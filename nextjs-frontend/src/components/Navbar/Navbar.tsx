import { AppBar, Box, ListItemText, Toolbar, Typography } from '@mui/material';
import AccountBalanceIcon from '@mui/icons-material/AccountBalance';
import { NavbarBankAccount } from './NavbarBankAccount';
import { cookies } from 'next/headers';

export function Navbar() {
  const cookiesStore = cookies();
  const bankAccountNumber = cookiesStore.get('bankAccountNumber')?.value;

  return (
    <AppBar position="fixed">
      <Toolbar sx={{ backgroundColor: 'primary' }}>
        <AccountBalanceIcon sx={{ mr: 1 }} />
        <Typography variant="h6" noWrap component="div">
          <ListItemText
            primary={process.env.NEXT_PUBLIC_BANK_NAME}
            secondary={`COD - ${process.env.NEXT_PUBLIC_BANK_CODE}`}
            secondaryTypographyProps={{ color: 'text.main' }}
          />
        </Typography>
        <Box flexGrow={1} />
        <Box>
          <div>
            {bankAccountNumber && (
              <NavbarBankAccount bankAccountNumber={bankAccountNumber} />
            )}
          </div>
        </Box>
      </Toolbar>
    </AppBar>
  );
}
