import { ExceptionFilter, Catch, ArgumentsHost } from '@nestjs/common';
import { Response } from 'express';
import { TransactionInvalidError } from '../transactions/transactions.service';

@Catch(TransactionInvalidError)
export class InvalidTransactionErrorFilter implements ExceptionFilter {
  catch(exception: Error, host: ArgumentsHost) {
    const ctx = host.switchToHttp();
    const response = ctx.getResponse<Response>();

    response.status(422).json({
      statusCode: 422,
      message: exception.message,
    });
  }
}
