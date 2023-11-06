'use client';
import {
  Alert,
  Box,
  Button,
  FormControl,
  FormControlLabel,
  FormLabel,
  Radio,
  RadioGroup,
  Snackbar,
  TextField,
  Typography,
} from '@mui/material';
import { useRouter } from 'next/navigation';
import { useState } from 'react';
import { createTransactionAction } from './create-transaction.action';
import { MyCard } from '../../../../components/MyCard';

export type WithdrawFormProps = {
  bankAccountNumber: string;
};

export function WithdrawForm(props: WithdrawFormProps) {
  const router = useRouter();
  const [open, setOpen] = useState(false);
  const createTransactionActionWithBankAccountNumber =
    createTransactionAction.bind(null, props.bankAccountNumber);

  function handleClose() {
    setOpen(false);
  }

  async function onSubmit(formData: FormData) {
    await createTransactionActionWithBankAccountNumber(formData);
    setOpen(true);
  }

  return (
    <div>
      <Typography variant="h5">Realizar transferência</Typography>
      <MyCard>
        <form
          style={{ display: 'flex', flexDirection: 'column' }}
          action={onSubmit}
        >
          <TextField
            name="bank_code_to"
            label="Código do banco do destinatário"
            margin="normal"
          />
          <TextField
            name="account_number_to"
            label="Número da conta do destinatário"
            margin="normal"
          />
          <FormControl sx={{ mt: 2 }} required>
            <FormLabel>Escolha um tipo de chave</FormLabel>
            <RadioGroup name="pix_key_kind_from">
              <FormControlLabel value="cpf" control={<Radio />} label="CPF" />
              <FormControlLabel
                value="email"
                control={<Radio />}
                label="Email"
              />
            </RadioGroup>
          </FormControl>
          <TextField
            name="pix_key_key_from"
            label="Digite sua chave pix"
            margin="normal"
          />
          <TextField
            name="amount"
            label="Valor"
            margin="normal"
            type="number"
          />
          <TextField name="description" label="Descrição" margin="normal" />
          <Box display={'flex'} gap={1} mt={2}>
            <Button type="submit" variant="contained">
              Concluir
            </Button>
            <Button
              type="button"
              variant="contained"
              color="secondary"
              onClick={() => {
                router.push(
                  `/bank-accounts/${props.bankAccountNumber}/dashboard`,
                );
              }}
            >
              Voltar
            </Button>
          </Box>
        </form>
      </MyCard>
      <Snackbar
        open={open}
        autoHideDuration={4000}
        onClose={handleClose}
        anchorOrigin={{
          vertical: 'top',
          horizontal: 'right',
        }}
      >
        <Alert onClose={handleClose} severity="success" sx={{ width: '100%' }}>
          Transferência agendada com sucesso!
        </Alert>
      </Snackbar>
    </div>
  );
}
