import { Box } from '@mui/material';
import { WithdrawForm } from './WithdrawForm';

export default async function WithdrawPage({
  params,
}: {
  params: { bankAccountNumber: string };
}) {
  return (
    <Box>
      <WithdrawForm bankAccountNumber={params.bankAccountNumber} />
    </Box>
  );
}
