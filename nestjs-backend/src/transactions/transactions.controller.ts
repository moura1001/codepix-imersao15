import {
  Controller,
  Get,
  Post,
  Body,
  Param,
  ValidationPipe,
  //  Patch,
  //  Param,
  //  Delete,
} from '@nestjs/common';
import { TransactionsService } from './transactions.service';
import { CreateTransactionDto } from './dto/create-transaction.dto';
import { MessagePattern, Payload } from '@nestjs/microservices';
import { ReceiveTransactionDto } from './dto/receive-transaction.dto';
import { TransactionStatus } from './entities/transaction.entity';
//import { UpdateTransactionDto } from './dto/update-transaction.dto';

@Controller('bank-accounts/:bankAccountNumberFrom/transactions')
export class TransactionsController {
  constructor(private readonly transactionsService: TransactionsService) {
    console.log('TransactionsController BANK_CODE -> ', process.env.BANK_CODE);
  }

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

  @MessagePattern('bank001')
  async onTransactionReceivedBank001(
    @Payload(new ValidationPipe()) message: ReceiveTransactionDto,
  ) {
    if (process.env.BANK_CODE !== '001') return;
    await this.processReceivedTransaction(message);
  }

  @MessagePattern('bank002')
  async onTransactionReceivedBank002(
    @Payload(new ValidationPipe()) message: ReceiveTransactionDto,
  ) {
    if (process.env.BANK_CODE !== '002') return;
    await this.processReceivedTransaction(message);
  }

  async processReceivedTransaction(
    @Payload(new ValidationPipe()) message: ReceiveTransactionDto,
  ) {
    console.log(process.env.BANK_CODE + ' RECEIVED MESSAGE ->', message);
    switch (message.status) {
      case TransactionStatus.PENDING:
        await this.transactionsService.receiveFromAnotherBank(message);
        return;
      case TransactionStatus.CONFIRMED:
        await this.transactionsService.completeTransaction(message);
        return;
    }
  }
}
