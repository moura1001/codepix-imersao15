import { Module } from '@nestjs/common';
import { ConfigModule } from '@nestjs/config';
import { AppController } from './app.controller';
import { AppService } from './app.service';
import { BankAccountsModule } from './bank-accounts/bank-accounts.module';
import { TypeOrmModule } from '@nestjs/typeorm';
import { BankAccount } from './bank-accounts/entities/bank-account.entity';

@Module({
  imports: [
    ConfigModule.forRoot(),
    TypeOrmModule.forRoot({
      type: 'postgres',
      url: (process.env.ENV_MODE === 'dev' || process.env.ENV_MODE === 'test') ? process.env.DSN_TEST : process.env.DSN,
      entities: [BankAccount],
      synchronize: /true/i.test(process.env.AUTO_MIGRATE_DB),
    }),
    BankAccountsModule,
  ],
  controllers: [AppController],
  providers: [AppService],
})
export class AppModule {}
