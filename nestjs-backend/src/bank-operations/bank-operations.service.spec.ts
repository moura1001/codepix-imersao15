import { Test, TestingModule } from '@nestjs/testing';
import { BankOperationsService } from './bank-operations.service';

describe('BankOperationsService', () => {
  let service: BankOperationsService;

  beforeEach(async () => {
    const module: TestingModule = await Test.createTestingModule({
      providers: [BankOperationsService],
    }).compile();

    service = module.get<BankOperationsService>(BankOperationsService);
  });

  it('should be defined', () => {
    expect(service).toBeDefined();
  });
});
