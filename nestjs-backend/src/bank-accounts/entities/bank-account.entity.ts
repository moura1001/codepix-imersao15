import {
  Column,
  Entity,
  PrimaryGeneratedColumn,
  CreateDateColumn,
} from 'typeorm';

@Entity({ name: 'bank_accounts' })
export class BankAccount {
  @PrimaryGeneratedColumn('uuid')
  account_number: string;

  @Column()
  owner_name: string;

  @Column({ default: 1000000000 })
  balance: number;

  @CreateDateColumn({ type: 'timestamp' })
  created_at: Date;
}
