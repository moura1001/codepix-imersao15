import Grid2 from '@mui/material/Unstable_Grid2/Grid2';
import { Typography } from '@mui/material';
import { CurrentBalance } from '../../../../components/CurrentBalance';
import { CardAction } from '@/components/CardAction';
import { MyLatestTransactions } from './MyLatestTransactions';
import { Transaction } from '../../../../models';

export async function getTransactions(
  bankAccountNumber: string,
): Promise<Transaction[]> {
  const response = await fetch(
    `http://localhost:3000/bank-accounts/${bankAccountNumber}/transactions`,
    {
      next: {
        revalidate: 20,
      },
    },
  );
  return response.json();
}

export default async function BankAccountDashboardPage({
  params,
}: {
  params: { bankAccountNumber: string };
}) {
  const transactions = await getTransactions(params.bankAccountNumber);

  return (
    <Grid2 container spacing={2}>
      <Grid2 xs={12} lg={6}>
        <div>
          <CurrentBalance bankAccountNumber={params.bankAccountNumber} />
        </div>
      </Grid2>
      <Grid2 container xs={12} lg={6} spacing={1}>
        <Grid2 xs={6}>
          <CardAction sx={{ display: 'flex', alignItems: 'center' }}>
            <Typography component="span" color={'primary'}>
              Transferência
            </Typography>
          </CardAction>
        </Grid2>
        <Grid2 xs={6}>
          <CardAction sx={{ display: 'flex', alignItems: 'center' }}>
            <Typography component="span" color={'primary'}>
              Nova chave pix
            </Typography>
          </CardAction>
        </Grid2>
      </Grid2>
      <Grid2 xs={12}>
        <Typography variant="h5">Últimos lançamentos</Typography>
        <MyLatestTransactions transactions={transactions} />
      </Grid2>
    </Grid2>
  );
}
