import { BankAccount } from '../models';

export async function getBankAccount(
  bankAccountNumber: string,
): Promise<BankAccount> {
  const response = await fetch(
    `${process.env.NEXT_PUBLIC_NEST_API_URL}/bank-accounts/${bankAccountNumber}`,
    {
      next: {
        revalidate: 20,
        tags: [`bank-accounts-${bankAccountNumber}`],
      },
    },
  );
  return response.json();
}
