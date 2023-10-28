import { Injectable, OnModuleInit } from '@nestjs/common';
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

@Injectable()
export class TransactionsService implements OnModuleInit {
  private BANK_CODE: string;

  constructor(
    @InjectRepository(Transaction)
    private transactionRepo: Repository<Transaction>,
    private configService: ConfigService,
    private dataSource: DataSource,
  ) {}

  onModuleInit() {
    this.BANK_CODE = this.configService.get('BANK_CODE');
  }

  // inicia a transferência
  async create(
    bankAccountNumberFrom: string,
    createTransactionDto: CreateTransactionDto,
  ) {
    return await this.dataSource.transaction(async (manager) => {
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

      return transaction;
    });
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
