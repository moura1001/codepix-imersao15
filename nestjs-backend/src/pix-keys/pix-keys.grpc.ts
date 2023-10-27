import { Observable } from 'rxjs';

export interface PixKeyRegistration {
  kind: string;
  key: string;
  accountNumber: string;
  bankCode: string;
}

export interface PixKey {
  kind: string;
  key: string;
}

export interface Bank {
  bankCode: string;
  bankName: string;
}

export interface Account {
  accountNumber: string;
  bankCode: string;
  bankName: string;
  ownerName: string;
  createdAt: string;
}

export interface PixKeyInfo {
  id: string;
  kind: string;
  key: string;
  account: Account;
  createdAt: string;
}

export interface RegisterCreatedResult {
  id: string;
  status: string;
  error: string;
}

export interface PixKeyClientGrpc {
  registerPixKey: (
    pixKey: PixKeyRegistration,
  ) => Observable<RegisterCreatedResult>;
  find: (pixKey: PixKey) => Observable<PixKeyInfo>;
  registerAccount: (account: Account) => Observable<RegisterCreatedResult>;
  registerBank: (bank: Bank) => Observable<RegisterCreatedResult>;
}
