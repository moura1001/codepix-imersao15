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
import { createPixKeyAction } from './create-pix-key.action';
import { MyCard } from '../../../../components/MyCard';

export type RegisterPixKeyFormProps = {
  bankAccountNumber: string;
};

export function RegisterPixKeyForm(props: RegisterPixKeyFormProps) {
  const router = useRouter();
  const [open, setOpen] = useState(false);
  const createPixKeyActionWithBankAccountNumber = createPixKeyAction.bind(
    null,
    props.bankAccountNumber,
  );

  function handleClose() {
    setOpen(false);
  }

  async function onSubmit(formData: FormData) {
    await createPixKeyActionWithBankAccountNumber(formData);
    setOpen(true);
  }

  return (
    <div>
      <Typography variant="h5">Cadastrar chave pix</Typography>
      <MyCard>
        <form
          style={{ display: 'flex', flexDirection: 'column' }}
          action={onSubmit}
        >
          <FormControl sx={{ mt: 2 }} required>
            <FormLabel>Escolha um tipo de chave</FormLabel>
            <RadioGroup name="kind">
              <FormControlLabel value="cpf" control={<Radio />} label="CPF" />
              <FormControlLabel
                value="email"
                control={<Radio />}
                label="Email"
              />
            </RadioGroup>
          </FormControl>
          <TextField name="key" label="Digite sua chave pix" margin="normal" />
          <Box display={'flex'} gap={1} mt={2}>
            <Button type="submit" variant="contained">
              Cadastrar
            </Button>
            <Button type="button" variant="contained" color="secondary">
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
          Chave pix cadastrada com sucesso!
        </Alert>
      </Snackbar>
    </div>
  );
}
