import { Inject, Injectable, OnModuleInit } from '@nestjs/common';
import { CreatePixKeyDto } from './dto/create-pix-key.dto';
//import { UpdatePixKeyDto } from './dto/update-pix-key.dto';
import { InjectRepository } from '@nestjs/typeorm';
import { PixKey, PixKeyKind } from './entities/pix-key.entity';
import { Repository } from 'typeorm';
import { BankAccount } from 'src/bank-accounts/entities/bank-account.entity';
import { ClientGrpc } from '@nestjs/microservices';
import {
  PixKeyClientGrpc,
  RegisterCreatedResult,
  PixKeyInfo,
  PixKeyRegistration,
  PixKey as PixKeyGrpcInput,
} from './pix-keys.grpc';
import { lastValueFrom } from 'rxjs';
import { ConfigService } from '@nestjs/config';

@Injectable()
export class PixKeysService implements OnModuleInit {
  private pixGrpcService: PixKeyClientGrpc;
  private BANK_CODE: string;

  constructor(
    @InjectRepository(PixKey)
    private pixKeyRepo: Repository<PixKey>,
    @InjectRepository(BankAccount)
    private bankAccountRepo: Repository<BankAccount>,
    @Inject('PIX_PACKAGE')
    private pixGrpcPackage: ClientGrpc,
    private configService: ConfigService,
  ) {}

  onModuleInit() {
    this.pixGrpcService = this.pixGrpcPackage.getService('PixService');
    this.BANK_CODE = this.configService.get('BANK_CODE');
  }

  async create(bankAccountNumber: string, createPixKeyDto: CreatePixKeyDto) {
    await this.bankAccountRepo.findOneOrFail({
      where: { account_number: bankAccountNumber },
    });
    // lógica para verificar se a chave pix existe no banco central (grpc)
    const remotePixKey = await this.findRemotePixKey(createPixKeyDto);
    if (remotePixKey && remotePixKey.id !== '') {
      return await this.createIfNotExists(remotePixKey);
    } else {
      const pixKey = {
        accountNumber: bankAccountNumber,
        bankCode: this.BANK_CODE,
        ...createPixKeyDto,
      } as PixKeyRegistration;

      console.log('pix key to create', pixKey);

      const remotePixKeyCreated = await this.registerRemotePixKey(pixKey);
      if (remotePixKeyCreated.status === 'created') {
        return this.pixKeyRepo.save({
          id: remotePixKeyCreated.id,
          bank_account_number: bankAccountNumber,
          ...createPixKeyDto,
        });
      } else {
        const errObj = {
          error: 'PixKeysService',
          method: 'create',
          grpcService: 'PixService.RegisterPixKey',
          response: remotePixKeyCreated,
        };
        throw new PixKeyRegisterGrpcError(JSON.stringify(errObj));
      }
    }
  }

  private async findRemotePixKey(
    pixKey: PixKeyGrpcInput,
  ): Promise<PixKeyInfo | null> {
    let errObj = {
      error: 'PixKeysService',
      method: 'findRemotePixKey',
    };
    try {
      return await lastValueFrom(this.pixGrpcService.find(pixKey));
    } catch (e) {
      errObj['details'] = e.details;
      console.log(errObj);
      return null;
    }

    errObj['details'] = 'gRPC unknow error';
    errObj['input'] = pixKey;
    throw new PixKeyUnknowGrpcError(JSON.stringify(errObj));
  }

  private async registerRemotePixKey(
    pixKey: PixKeyRegistration,
  ): Promise<RegisterCreatedResult> {
    try {
      return await lastValueFrom(this.pixGrpcService.registerPixKey(pixKey));
    } catch (e) {
      const errObj = {
        error: 'PixKeysService',
        method: 'registerRemotePixKey',
        details: e.details,
      };
      console.log(errObj);
      throw new PixKeyRegisterGrpcError(JSON.stringify(errObj));
    }
  }

  private async createIfNotExists(remotePixKey: PixKeyInfo) {
    const hasLocalPixKey = await this.pixKeyRepo.exist({
      where: {
        id: remotePixKey.id,
        key: remotePixKey.key,
        kind: remotePixKey.kind as PixKeyKind,
        bank_account_number: remotePixKey.account.accountNumber,
      },
    });

    if (hasLocalPixKey) {
      throw new PixKeyAlreadyExistsError('pix key already exist');
    } else {
      return this.pixKeyRepo.save({
        id: remotePixKey.id,
        bank_account_number: remotePixKey.account.accountNumber,
        account: remotePixKey.account,
        key: remotePixKey.key,
        kind: remotePixKey.kind as PixKeyKind,
        created_at: remotePixKey.createdAt,
      });
    }
  }

  async findAll(bankAccountNumber: string) {
    return await this.pixKeyRepo.find({
      where: { bank_account_number: bankAccountNumber },
      order: { created_at: 'DESC' },
    });
  }

  /*findOne(id: number) {
    return `This action returns a #${id} pixKey`;
  }

  update(id: number, updatePixKeyDto: UpdatePixKeyDto) {
    return `This action updates a #${id} pixKey`;
  }

  remove(id: number) {
    return `This action removes a #${id} pixKey`;
  }*/
}

export class PixKeyUnknowGrpcError extends Error {}
export class PixKeyRegisterGrpcError extends Error {}
export class PixKeyAlreadyExistsError extends Error {}
