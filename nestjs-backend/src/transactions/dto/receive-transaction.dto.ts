import {
  IsIn,
  IsNotEmpty,
  IsOptional,
  IsString,
  IsUUID,
  IsNumber,
  Min,
} from 'class-validator';
import { PixKeyKind } from 'src/pix-keys/entities/pix-key.entity';
import { TransactionStatus } from '../entities/transaction.entity';

export class ReceiveTransactionDto {
  @IsUUID()
  id: string;

  @IsIn([
    TransactionStatus.CANCELLED,
    TransactionStatus.COMPLETED,
    TransactionStatus.CONFIRMED,
    TransactionStatus.PENDING,
  ])
  @IsString()
  @IsNotEmpty()
  status: TransactionStatus;

  @IsString()
  @IsOptional()
  cancelDescription: string | null = null;

  @IsUUID()
  relatedTransactionIdFrom: string;

  @IsString()
  @IsNotEmpty()
  bankCodeFrom: string;

  @IsString()
  @IsNotEmpty()
  bankCodeTo: string;

  @IsUUID()
  accountNumberTo: string;

  @IsNumber({ maxDecimalPlaces: 2 })
  @Min(1)
  @IsNotEmpty()
  amount: number;

  @IsString()
  @IsNotEmpty()
  pixKeyFrom: string;

  @IsIn([PixKeyKind.cpf, PixKeyKind.email])
  @IsString()
  @IsNotEmpty()
  pixKeyFromKind: PixKeyKind;

  @IsString()
  @IsOptional()
  description: string | null = null;
}
