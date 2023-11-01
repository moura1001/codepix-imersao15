import { NestFactory } from '@nestjs/core';
import { AppModule } from './app.module';
import { GrpcUnknowErrorFilter } from './exceptions-handlers/grpc-unknow.error';
import { RegisterErrorFilter } from './exceptions-handlers/register.error';
import { NotFoundErrorFilter } from './exceptions-handlers/not-found.error';
import { ValidationPipe } from '@nestjs/common';
import { InvalidTransactionErrorFilter } from './exceptions-handlers/invalid-transaction.error';
import { KafkaUnknowErrorFilter } from './exceptions-handlers/kafka-unknow.error';
import { ConfigService } from '@nestjs/config';
import { Transport } from '@nestjs/microservices';

async function bootstrap() {
  const app = await NestFactory.create(AppModule);

  const configService = app.get<ConfigService>(ConfigService);

  app.useGlobalFilters(
    new GrpcUnknowErrorFilter(),
    new RegisterErrorFilter(),
    new NotFoundErrorFilter(),
    new InvalidTransactionErrorFilter(),
    new KafkaUnknowErrorFilter(),
  );

  app.useGlobalPipes(
    new ValidationPipe({
      errorHttpStatusCode: 422,
    }),
  );

  app.connectMicroservice({
    transport: Transport.KAFKA,
    options: {
      client: {
        brokers: [configService.get('KAFKA_BOOTSTRAP_SERVERS')],
      },
      consumer: {
        groupId: configService.get('KAFKA_CONSUMER_GROUP_ID'),
      },
    },
  });

  await app.startAllMicroservices();
  await app.listen(configService.get('PORT') || 3000);
}
bootstrap();
