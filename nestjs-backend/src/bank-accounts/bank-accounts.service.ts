import { Inject, Injectable, OnModuleInit } from '@nestjs/common';
import { CreateBankAccountDto } from './dto/create-bank-account.dto';
//import { UpdateBankAccountDto } from './dto/update-bank-account.dto';
import { Repository } from 'typeorm';
import { BankAccount } from './entities/bank-account.entity';
import { InjectRepository } from '@nestjs/typeorm';
import {
  PixKeyClientGrpc,
  Account as AccountGrpcInput,
  RegisterCreatedResult,
} from 'src/pix-keys/pix-keys.grpc';
import { ClientGrpc } from '@nestjs/microservices';
import { ConfigService } from '@nestjs/config';
import { lastValueFrom } from 'rxjs';

@Injectable()
export class BankAccountsService implements OnModuleInit {
  private pixGrpcService: PixKeyClientGrpc;
  private BANK_CODE: string;

  constructor(
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

  async create(createBankAccountDto: CreateBankAccountDto) {
    try {
      const hasAccount = await this.bankAccountRepo.exist({
        where: {
          account_number: createBankAccountDto.account_number,
        },
      });
      if (!hasAccount) {
        return this.bankAccountRepo.save(createBankAccountDto);
      }
      throw new AccountAlreadyExistsError(
        'account ' +
          JSON.stringify(createBankAccountDto) +
          ' already registrated',
      );
    } catch (e) {
      const errObj = {
        error: 'BankAccountsService',
        method: 'create',
        details: e,
      };
      console.log(errObj);
      throw e;
    }
  }

  async registerOnPix(createBankAccountDto: CreateBankAccountDto) {
    await this.findOne(createBankAccountDto.account_number);

    const account = {
      bankCode: this.BANK_CODE,
      accountNumber: createBankAccountDto.account_number,
      ownerName: createBankAccountDto.owner_name,
    } as AccountGrpcInput;

    const remotePixAccountRegistred = await this.registerPixAccount(account);
    if (remotePixAccountRegistred.status === 'created') {
      return createBankAccountDto;
    } else {
      const errObj = {
        error: 'BankAccountsService',
        method: 'registerOnPix',
        grpcService: 'PixService.RegisterAccount',
        response: remotePixAccountRegistred,
      };
      console.log(errObj);
      throw new PixAccountRegisterGrpcError(JSON.stringify(errObj));
    }
  }

  private async registerPixAccount(
    account: AccountGrpcInput,
  ): Promise<RegisterCreatedResult> {
    try {
      return await lastValueFrom(this.pixGrpcService.registerAccount(account));
    } catch (e) {
      const errObj = {
        error: 'BankAccountsService',
        method: 'registerPixAccount',
        details: e.details,
      };
      console.log(errObj);
      throw new PixAccountRegisterUnknowGrpcError(JSON.stringify(errObj));
    }
  }

  findAll() {
    return this.bankAccountRepo.find();
  }

  async findOne(account_number: string) {
    try {
      return await this.bankAccountRepo.findOneOrFail({
        where: { account_number },
      });
    } catch (e) {
      throw new AccountNotFoundError(
        'account ' + account_number + ' not found',
      );
    }
  }

  async save(bankAccount: BankAccount) {
    return await this.bankAccountRepo.save(bankAccount);
  }

  /*update(id: number, updateBankAccountDto: UpdateBankAccountDto) {
    return `This action updates a #${id} bankAccount`;
  }

  remove(id: number) {
    return `This action removes a #${id} bankAccount`;
  }*/
}

export class AccountAlreadyExistsError extends Error {}
export class PixAccountRegisterGrpcError extends Error {}
export class PixAccountRegisterUnknowGrpcError extends Error {}
export class AccountNotFoundError extends Error {}
