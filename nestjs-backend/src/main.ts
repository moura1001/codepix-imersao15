import { NestFactory } from '@nestjs/core';
import { AppModule } from './app.module';
import { GrpcUnknowErrorFilter } from './exceptions-handlers/grpc-unknow.error';
import { RegisterErrorFilter } from './exceptions-handlers/register.error';
import { NotFoundErrorFilter } from './exceptions-handlers/not-found.error';
import { ValidationPipe } from '@nestjs/common';
import { InvalidTransactionErrorFilter } from './exceptions-handlers/invalid-transaction.error';

async function bootstrap() {
  const app = await NestFactory.create(AppModule);

  app.useGlobalFilters(
    new GrpcUnknowErrorFilter(),
    new RegisterErrorFilter(),
    new NotFoundErrorFilter(),
    new InvalidTransactionErrorFilter(),
  );

  app.useGlobalPipes(
    new ValidationPipe({
      errorHttpStatusCode: 422,
    }),
  );

  await app.listen(3000);
}
bootstrap();
