import { NestFactory } from '@nestjs/core';
import { AppModule } from './app.module';
import { GrpcUnknowErrorFilter } from './exceptions-handlers/grpc-unknow.error';
import { RegisterErrorFilter } from './exceptions-handlers/register.error';
import { NotFoundErrorFilter } from './exceptions-handlers/not-found.error';

async function bootstrap() {
  const app = await NestFactory.create(AppModule);

  app.useGlobalFilters(
    new GrpcUnknowErrorFilter(),
    new RegisterErrorFilter(),
    new NotFoundErrorFilter(),
  );

  await app.listen(3000);
}
bootstrap();
