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

export class CreateTransactionDto {
  @IsString()
  @IsNotEmpty()
  bank_code_to: string;

  @IsUUID()
  account_number_to: string;

  @IsString()
  @IsNotEmpty()
  pix_key_key_from: string;

  @IsIn([PixKeyKind.cpf, PixKeyKind.email])
  @IsString()
  @IsNotEmpty()
  pix_key_kind_from: PixKeyKind;

  @IsString()
  @IsOptional()
  description: string | null = null;

  @IsNumber({ maxDecimalPlaces: 2 })
  @Min(1)
  @IsNotEmpty()
  amount: number;
}
