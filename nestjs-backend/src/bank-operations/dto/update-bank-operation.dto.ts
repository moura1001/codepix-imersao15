import { PartialType } from '@nestjs/mapped-types';
import { CreateBankOperationDto } from './create-bank-operation.dto';

export class UpdateBankOperationDto extends PartialType(CreateBankOperationDto) {}
