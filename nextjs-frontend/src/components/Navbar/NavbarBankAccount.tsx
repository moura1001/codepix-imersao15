import { getBankAccount } from '@/service/get-bank-account.query';
import { Box, Chip, Button } from '@mui/material';
import PersonIcon from '@mui/icons-material/Person';
import { cookies } from 'next/headers';
import { redirect } from 'next/navigation';

export type NavbarBankAccountProps = {
  bankAccountNumber: string;
};

export async function NavbarBankAccount(props: NavbarBankAccountProps) {
  const bankAccount = await getBankAccount(props.bankAccountNumber);
  return (
    <Box>
      <Chip
        label={
          <Box>
            <PersonIcon />
            <Box display={'flex'} flexDirection={'column'}>
              <span>{bankAccount.owner_name}</span>
              <span>C/C {bankAccount.account_number}</span>
            </Box>
          </Box>
        }
        sx={{ backgroundColor: 'primary.contrastText', py: 3 }}
      />
      <form
        action={async () => {
          'use server';
          cookies().delete('bankAccountNumber');
          redirect('/');
        }}
        style={{ all: 'unset' }}
      >
        <Button color="inherit" type="submit">
          Sair
        </Button>
      </form>
    </Box>
  );
}
