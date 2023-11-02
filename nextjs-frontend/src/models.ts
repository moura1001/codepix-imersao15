export interface BankAccount {
  account_number: string;
  owner_name: string;
  balance: number;
  created_at: Date;
}

export enum PixKeyKind {
  cpf = 'cpf',
  email = 'email',
}

export interface PixKey {
  id: string;
  kind: PixKeyKind;
  key: string;
  bank_account_number: string;
  bankAccount: BankAccount;
  created_at: Date;
}

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

export interface Transaction {
  id: string;
  related_transaction_id: string;
  amount: number;
  description: string;
  bank_code_from: string;
  account_number_from: string;
  bank_code_to: string;
  account_number_to: string;
  pix_key_key_from: string;
  pix_key_kind_from: PixKeyKind;
  status: TransactionStatus = TransactionStatus.PENDING;
  operation: TransactionOperation;
  created_at: Date;
}
