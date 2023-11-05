import { Box } from '@mui/material';
import Grid2 from '@mui/material/Unstable_Grid2/Grid2';
import { PixKeyList } from './PixKeyList';
import { RegisterPixKeyForm } from './RegisterPixKeyForm';

export default function MyPixDashboardPage({
  params,
}: {
  params: { bankAccountNumber: string };
}) {
  return (
    <Box>
      <Grid2 container spacing={8}>
        <Grid2 xs={12} sm={6}>
          <RegisterPixKeyForm bankAccountNumber={params.bankAccountNumber} />
        </Grid2>

        <Grid2 xs={12} sm={6}>
          <div>
            <PixKeyList bankAccountNumber={params.bankAccountNumber} />
          </div>
        </Grid2>
      </Grid2>
    </Box>
  );
}
