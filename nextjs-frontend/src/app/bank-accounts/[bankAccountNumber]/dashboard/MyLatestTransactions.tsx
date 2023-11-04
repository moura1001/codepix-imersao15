'use client';
import { DataGrid, GridColDef } from '../../../../components/mui';
import { red, green } from '@mui/material/colors';
import { Transaction } from '../../../../models';

const columns: GridColDef[] = [
  { field: 'id', headerName: 'ID', width: 300 },
  {
    field: 'created_at',
    headerName: 'Data',
    width: 200,
    renderCell: (params) => new Date(params.value as string).toLocaleString(),
  },
  { field: 'description', headerName: 'Descrição', width: 130 },
  {
    field: 'amount',
    headerName: 'Valor (R$)',
    width: 180,
    renderCell: (params) => {
      const amount = new Intl.NumberFormat('pt-BR', {
        style: 'currency',
        currency: 'BRL',
      }).format(params.value as number);
      return (
        <span style={{ color: params.value < 0 ? red[500] : green[500] }}>
          {amount}
        </span>
      );
    },
  },
];

export type MyLatestTransactionsProps = {
	transactions: Transaction[]
}

export function MyLatestTransactions(props: MyLatestTransactionsProps) {
  return <DataGrid rows={props.transactions} columns={columns} />;
}
