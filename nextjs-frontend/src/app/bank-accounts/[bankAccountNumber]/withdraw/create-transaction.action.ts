'use server';
import { revalidateTag } from 'next/cache';
import { redirect } from 'next/navigation';

export async function createTransactionAction(
  bankAccountNumber: string,
  formData: FormData,
) {
  const bankCodeTo = formData.get('bank_code_to');
  const accountNumberTo = formData.get('account_number_to');
  const pixKeyKindFrom = formData.get('pix_key_kind_from');
  const pixKeyKeyFrom = formData.get('pix_key_key_from');
  const amount: number = +formData.get('amount');
  const description = formData.get('description');

  const response = await fetch(
    `http://localhost:3000/bank-accounts/${bankAccountNumber}/transactions`,
    {
      method: 'POST',
      body: JSON.stringify({
        bank_code_to: bankCodeTo,
        account_number_to: accountNumberTo,
        pix_key_key_from: pixKeyKeyFrom,
        pix_key_kind_from: pixKeyKindFrom,
        amount: amount,
        description: description,
      }),
      headers: {
        'Content-Type': 'application/json',
      },
    },
  );

  if (!response.ok) {
    throw new Error(JSON.stringify(await response.json()));
  }

  //const result = await response.json();
  revalidateTag(`bank-accounts-${bankAccountNumber}`);
  redirect(`/bank-accounts/${bankAccountNumber}/dashboard`);
}
