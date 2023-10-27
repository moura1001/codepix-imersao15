import {
  Controller,
  //  Get,
  Post,
  //  Body,
  //  Patch,
  //  Param,
  //  Delete,
} from '@nestjs/common';
import { BankOperationsService } from './bank-operations.service';
//import { CreateBankOperationDto } from './dto/create-bank-operation.dto';
//import { UpdateBankOperationDto } from './dto/update-bank-operation.dto';

@Controller('bank-operations')
export class BankOperationsController {
  constructor(private readonly bankOperationsService: BankOperationsService) {}

  @Post('pix-system')
  registerOnPixSystem() {
    return this.bankOperationsService.registerOnPix();
  }

  /*@Post()
  create(@Body() createBankOperationDto: CreateBankOperationDto) {
    return this.bankOperationsService.create(createBankOperationDto);
  }

  @Get()
  findAll() {
    return this.bankOperationsService.findAll();
  }

  @Get(':id')
  findOne(@Param('id') id: string) {
    return this.bankOperationsService.findOne(+id);
  }

  @Patch(':id')
  update(@Param('id') id: string, @Body() updateBankOperationDto: UpdateBankOperationDto) {
    return this.bankOperationsService.update(+id, updateBankOperationDto);
  }

  @Delete(':id')
  remove(@Param('id') id: string) {
    return this.bankOperationsService.remove(+id);
  }*/
}
