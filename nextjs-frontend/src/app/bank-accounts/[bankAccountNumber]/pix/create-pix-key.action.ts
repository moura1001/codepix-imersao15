'use server';
import { revalidateTag } from 'next/cache';

export async function createPixKeyAction(
  bankAccountNumber: string,
  formData: FormData,
) {
  const pixKeyKind = formData.get('kind');
  const pixKeyKey = formData.get('key');
  const response = await fetch(
    `http://localhost:3000/bank-accounts/${bankAccountNumber}/pix-keys`,
    {
      method: 'POST',
      body: JSON.stringify({
        kind: pixKeyKind,
        key: pixKeyKey,
      }),
      headers: {
        'Content-Type': 'application/json',
      },
    },
  );

  if (!response.ok) {
    throw new Error(JSON.stringify(await response.json()));
  }

  const result = await response.json();
  revalidateTag(`pix-keys-${bankAccountNumber}`);
  return result;
}
