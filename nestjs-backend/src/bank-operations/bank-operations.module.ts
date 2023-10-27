import { Module } from '@nestjs/common';
import { BankOperationsService } from './bank-operations.service';
import { BankOperationsController } from './bank-operations.controller';
import { ClientsModule, Transport } from '@nestjs/microservices';
import { join } from 'path';
import { ConfigService, ConfigModule } from '@nestjs/config';

@Module({
  imports: [
    ClientsModule.registerAsync([
      {
        name: 'PIX_PACKAGE',
        useFactory: (configService: ConfigService) => ({
          transport: Transport.GRPC,
          options: {
            url: configService.get('GRPC_URL'),
            package: configService.get('GRPC_PACKAGE'),
            protoPath: join(__dirname, '..', 'protofiles', 'pixKey.proto'),
          },
        }),
        inject: [ConfigService],
      },
    ]),
    ConfigModule,
  ],
  controllers: [BankOperationsController],
  providers: [BankOperationsService],
})
export class BankOperationsModule {}
