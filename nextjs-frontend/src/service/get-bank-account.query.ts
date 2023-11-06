import { BankAccount } from '../models';

export async function getBankAccount(
  bankAccountNumber: string,
): Promise<BankAccount> {
  const response = await fetch(
    `http://localhost:3000/bank-accounts/${bankAccountNumber}`,
    {
      next: {
        revalidate: 20,
        tags: [`bank-accounts-${bankAccountNumber}`],
      },
    },
  );
  return response.json();
}
