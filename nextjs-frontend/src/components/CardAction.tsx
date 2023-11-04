import Grid2 from '@mui/material/Unstable_Grid2/Grid2';
import { Button, SxProps, Theme } from '@mui/material';
import { ArrowForwardIos } from '@mui/icons-material';
import { PropsWithChildren } from 'react';
import { MyCard } from './MyCard';

export type CardActionProps = {
  sx?: SxProps<Theme>;
};

export function CardAction(props: PropsWithChildren<CardActionProps>) {
  return (
    <MyCard>
      <Grid2 container>
        <Grid2 xs={12} sm={9} sx={props.sx}>
          {props.children}
        </Grid2>
        <Grid2
          xs={12}
          sm={3}
          display={'flex'}
          alignItems={'center'}
          justifyContent={'flex-end'}
        >
          <Button color="primary">
            <ArrowForwardIos />
          </Button>
        </Grid2>
      </Grid2>
    </MyCard>
  );
}
