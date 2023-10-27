import { Test, TestingModule } from '@nestjs/testing';
import { BankOperationsController } from './bank-operations.controller';
import { BankOperationsService } from './bank-operations.service';

describe('BankOperationsController', () => {
  let controller: BankOperationsController;

  beforeEach(async () => {
    const module: TestingModule = await Test.createTestingModule({
      controllers: [BankOperationsController],
      providers: [BankOperationsService],
    }).compile();

    controller = module.get<BankOperationsController>(BankOperationsController);
  });

  it('should be defined', () => {
    expect(controller).toBeDefined();
  });
});
