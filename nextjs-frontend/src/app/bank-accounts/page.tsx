import { BankAccount } from '../../models';

export async function getBankAccounts(): Promise<BankAccount[]> {
  const response = await fetch('http://localhost:3000/bank-accounts');
  return response.json();
}

export default async function HomePage() {
  const bankAccounts = await getBankAccounts();
  return (
    <div>
      <h1>Home Page</h1>
      <ul>
        {bankAccounts.map((bankAccount) => (
          <li key={bankAccount.account_number}>
            {bankAccount.owner_name} - {bankAccount.balance}
          </li>
        ))}
      </ul>
    </div>
  );
}
