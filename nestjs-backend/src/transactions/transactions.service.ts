import { Inject, Injectable, OnModuleInit } from '@nestjs/common';
import { DataSource, Repository } from 'typeorm';
import {
  Transaction,
  TransactionOperation,
} from './entities/transaction.entity';
import { InjectRepository } from '@nestjs/typeorm';
import { CreateTransactionDto } from './dto/create-transaction.dto';
//import { UpdateTransactionDto } from './dto/update-transaction.dto';
import { BankAccount } from '../bank-accounts/entities/bank-account.entity';
import { ConfigService } from '@nestjs/config';
import { AccountNotFoundError } from 'src/bank-accounts/bank-accounts.service';
import { ClientKafka } from '@nestjs/microservices';
import { lastValueFrom } from 'rxjs';

@Injectable()
export class TransactionsService implements OnModuleInit {
  private BANK_CODE: string;
  private TRANSACTIONS_TOPIC: string;
  private TRANSACTION_CONFIRMATION_TOPIC: string;

  constructor(
    @InjectRepository(Transaction)
    private transactionRepo: Repository<Transaction>,
    private configService: ConfigService,
    private dataSource: DataSource,
    @Inject('KAFKA_SERVICE')
    private kafkaService: ClientKafka,
  ) {}

  onModuleInit() {
    this.BANK_CODE = this.configService.get('BANK_CODE');
    this.TRANSACTIONS_TOPIC = this.configService.get('TRANSACTIONS_TOPIC');
    this.TRANSACTION_CONFIRMATION_TOPIC = this.configService.get(
      'TRANSACTION_CONFIRMATION_TOPIC',
    );
  }

  // inicia a transferência
  async create(
    bankAccountNumberFrom: string,
    createTransactionDto: CreateTransactionDto,
  ) {
    const transaction = await this.dataSource.transaction(async (manager) => {
      let bankAccount: BankAccount;
      try {
        bankAccount = await manager.findOneOrFail(BankAccount, {
          where: { account_number: bankAccountNumberFrom },
          lock: { mode: 'pessimistic_write' },
        });
      } catch (e) {
        throw new AccountNotFoundError(
          'account ' + bankAccountNumberFrom + ' does not exist',
        );
      }

      bankAccount.balance -= createTransactionDto.amount;
      if (bankAccount.balance < 0) {
        throw new TransactionInvalidError('insufficient funds');
      }

      const transaction = manager.create(Transaction, {
        ...createTransactionDto,
        amount: createTransactionDto.amount * -1,
        bank_code_from: this.BANK_CODE,
        account_number_from: bankAccountNumberFrom,
        operation: TransactionOperation.DEBIT,
      });

      await manager.save<Transaction>(transaction);

      await manager.save<BankAccount>(bankAccount);

      const sendData = {
        relatedTransactionIdFrom: transaction.id,
	      bankCodeFrom: this.BANK_CODE,
        bankCodeTo: createTransactionDto.bank_code_to,
        accountNumberTo: createTransactionDto.account_number_to,
        amount: createTransactionDto.amount,
        pixKeyFrom: createTransactionDto.pix_key_key_from,
        pixKeyFromKind: createTransactionDto.pix_key_kind_from,
        description: createTransactionDto.description,
      };

      try {
        await lastValueFrom(
          this.kafkaService.emit(this.TRANSACTIONS_TOPIC, sendData),
        );
      } catch (e) {
        const errObj = {
          error: 'TransactionsService',
          method: 'create',
          details: e.details,
        };
        console.log(errObj);
        throw new TransactionUnknowKafkaError(JSON.stringify(errObj));
      }

      return transaction;
    });

    return transaction;
  }

  findAll(bankAccountNumber: string) {
    return this.transactionRepo.find({
      where: { account_number_from: bankAccountNumber },
      order: { created_at: 'DESC' },
    });
  }

  /*findOne(id: number) {
    return `This action returns a #${id} transaction`;
  }

  update(id: number, updateTransactionDto: UpdateTransactionDto) {
    return `This action updates a #${id} transaction`;
  }

  remove(id: number) {
    return `This action removes a #${id} transaction`;
  }*/
}

export class TransactionInvalidError extends Error {}
export class TransactionUnknowKafkaError extends Error {}
