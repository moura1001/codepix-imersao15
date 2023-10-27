import { ExceptionFilter, Catch, ArgumentsHost } from '@nestjs/common';
import {
  PixKeyAlreadyExistsError,
  PixKeyRegisterGrpcError,
} from 'src/pix-keys/pix-keys.service';
import { Response } from 'express';
import {
  AccountAlreadyExistsError,
  PixAccountRegisterGrpcError,
} from 'src/bank-accounts/bank-accounts.service';
import { PixBankRegisterGrpcError } from 'src/bank-operations/bank-operations.service';

@Catch(
  PixKeyAlreadyExistsError,
  AccountAlreadyExistsError,
  PixKeyRegisterGrpcError,
  PixAccountRegisterGrpcError,
  PixBankRegisterGrpcError,
)
export class RegisterErrorFilter implements ExceptionFilter {
  catch(exception: Error, host: ArgumentsHost) {
    const ctx = host.switchToHttp();
    const response = ctx.getResponse<Response>();

    response.status(422).json({
      statusCode: 422,
      message: exception.message,
    });
  }
}
