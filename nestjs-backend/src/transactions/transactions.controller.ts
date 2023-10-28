import {
  Controller,
  Get,
  Post,
  Body,
  Param,
  //  Patch,
  //  Param,
  //  Delete,
} from '@nestjs/common';
import { TransactionsService } from './transactions.service';
import { CreateTransactionDto } from './dto/create-transaction.dto';
//import { UpdateTransactionDto } from './dto/update-transaction.dto';

@Controller('bank-accounts/:bankAccountNumberFrom/transactions')
export class TransactionsController {
  constructor(private readonly transactionsService: TransactionsService) {}

  @Post()
  create(
    @Param('bankAccountNumberFrom') bankAccountNumberFrom: string,
    @Body() createTransactionDto: CreateTransactionDto,
  ) {
    return this.transactionsService.create(
      bankAccountNumberFrom,
      createTransactionDto,
    );
  }

  @Get()
  findAll(@Param('bankAccountNumberFrom') bankAccountNumberFrom: string) {
    return this.transactionsService.findAll(bankAccountNumberFrom);
  }

  /*@Get(':id')
  findOne(@Param('id') id: string) {
    return this.transactionsService.findOne(+id);
  }

  @Patch(':id')
  update(@Param('id') id: string, @Body() updateTransactionDto: UpdateTransactionDto) {
    return this.transactionsService.update(+id, updateTransactionDto);
  }

  @Delete(':id')
  remove(@Param('id') id: string) {
    return this.transactionsService.remove(+id);
  }*/
}
