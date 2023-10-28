import { Module } from '@nestjs/common';
import { ConfigModule, ConfigService } from '@nestjs/config';
import { AppController } from './app.controller';
import { AppService } from './app.service';
import { BankAccountsModule } from './bank-accounts/bank-accounts.module';
import { TypeOrmModule } from '@nestjs/typeorm';
import { BankAccount } from './bank-accounts/entities/bank-account.entity';
import { PixKeysModule } from './pix-keys/pix-keys.module';
import { PixKey } from './pix-keys/entities/pix-key.entity';
import { BankOperationsModule } from './bank-operations/bank-operations.module';
import { TransactionsModule } from './transactions/transactions.module';
import { Transaction } from './transactions/entities/transaction.entity';

@Module({
  imports: [
    ConfigModule.forRoot({
      envFilePath: ['.env'],
      isGlobal: true,
    }),
    TypeOrmModule.forRootAsync({
      useFactory: (configService: ConfigService) => ({
        type: 'postgres',
        url:
          configService.get('ENV_MODE') === 'dev' ||
          configService.get('ENV_MODE') === 'test'
            ? configService.get('DSN_TEST')
            : configService.get('DSN'),
        entities: [BankAccount, PixKey, Transaction],
        synchronize: /true/i.test(configService.get('AUTO_MIGRATE_DB')),
        logging: /true/i.test(configService.get('AUTO_MIGRATE_DB')),
      }),
      inject: [ConfigService],
    }),
    BankAccountsModule,
    PixKeysModule,
    BankOperationsModule,
    TransactionsModule,
  ],
  controllers: [AppController],
  providers: [AppService],
})
export class AppModule {}
