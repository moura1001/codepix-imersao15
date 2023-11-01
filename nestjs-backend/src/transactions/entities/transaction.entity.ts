import { BankAccount } from 'src/bank-accounts/entities/bank-account.entity';
import { PixKeyKind } from 'src/pix-keys/entities/pix-key.entity';
import {
  Column,
  CreateDateColumn,
  Entity,
  JoinColumn,
  ManyToOne,
  PrimaryGeneratedColumn,
} from 'typeorm';

export enum TransactionStatus {
  PENDING = 'pending',
  CANCELLED = 'cancelled',
  CONFIRMED = 'comfirmed',
  COMPLETED = 'completed',
}

export enum TransactionOperation {
  DEBIT = 'debit',
  CREDIT = 'credit',
}

@Entity({ name: 'transactions' })
export class Transaction {
  @PrimaryGeneratedColumn('uuid')
  id: string;

  @Column({ nullable: true, type: 'uuid' })
  related_transaction_id: string;

  @Column()
  amount: number;

  @Column({ nullable: true })
  description: string;

  @ManyToOne(() => BankAccount)
  @JoinColumn({ name: 'account_number_from' })
  bankAccount: BankAccount;

  @Column()
  bank_code_from: string;

  @Column({ nullable: true })
  account_number_from: string;

  @Column()
  bank_code_to: string;

  @Column({ type: 'uuid' })
  account_number_to: string;

  @Column()
  pix_key_key_from: string;

  @Column()
  pix_key_kind_from: PixKeyKind;

  @Column()
  status: TransactionStatus = TransactionStatus.PENDING;

  @Column()
  operation: TransactionOperation;

  @CreateDateColumn({ type: 'timestamp' })
  created_at: Date;
}
