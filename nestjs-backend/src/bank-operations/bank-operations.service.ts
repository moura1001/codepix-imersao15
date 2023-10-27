import { Inject, Injectable, OnModuleInit } from '@nestjs/common';
//import { CreateBankOperationDto } from './dto/create-bank-operation.dto';
//import { UpdateBankOperationDto } from './dto/update-bank-operation.dto';
import {
  PixKeyClientGrpc,
  Bank as BankGrpcInput,
  RegisterCreatedResult,
} from 'src/pix-keys/pix-keys.grpc';
import { ClientGrpc } from '@nestjs/microservices';
import { ConfigService } from '@nestjs/config';
import { lastValueFrom } from 'rxjs';

@Injectable()
export class BankOperationsService implements OnModuleInit {
  private pixGrpcService: PixKeyClientGrpc;
  private BANK_CODE: string;
  private BANK_NAME: string;

  constructor(
    @Inject('PIX_PACKAGE')
    private pixGrpcPackage: ClientGrpc,
    private configService: ConfigService,
  ) {}

  onModuleInit() {
    this.pixGrpcService = this.pixGrpcPackage.getService('PixService');
    this.BANK_CODE = this.configService.get('BANK_CODE');
    this.BANK_NAME = this.configService.get('BANK_NAME');
  }

  async registerOnPix() {
    const bank = {
      bankCode: this.BANK_CODE,
      bankName: this.BANK_NAME,
    } as BankGrpcInput;

    const remotePixBankRegistred = await this.registerPixBank(bank);
    if (remotePixBankRegistred.status === 'created') {
      return bank;
    } else {
      const errObj = {
        error: 'BankOperationsService',
        method: 'registerOnPix',
        grpcService: 'PixService.RegisterBank',
        response: remotePixBankRegistred,
      };
      console.log(errObj);
      throw new PixBankRegisterGrpcError(JSON.stringify(errObj));
    }
  }

  private async registerPixBank(
    bank: BankGrpcInput,
  ): Promise<RegisterCreatedResult> {
    try {
      return await lastValueFrom(this.pixGrpcService.registerBank(bank));
    } catch (e) {
      const errObj = {
        error: 'BankOperationsService',
        method: 'registerPixBank',
        details: e.details,
      };
      console.log(errObj);
      throw new PixBankRegisterUnknowGrpcError(JSON.stringify(errObj));
    }
  }

  /*create(createBankOperationDto: CreateBankOperationDto) {
    return 'This action adds a new bankOperation';
  }

  findAll() {
    return `This action returns all bankOperations`;
  }

  findOne(id: number) {
    return `This action returns a #${id} bankOperation`;
  }

  update(id: number, updateBankOperationDto: UpdateBankOperationDto) {
    return `This action updates a #${id} bankOperation`;
  }

  remove(id: number) {
    return `This action removes a #${id} bankOperation`;
  }*/
}

export class PixBankRegisterGrpcError extends Error {}
export class PixBankRegisterUnknowGrpcError extends Error {}
