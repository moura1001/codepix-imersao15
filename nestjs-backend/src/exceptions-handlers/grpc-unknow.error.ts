import { ExceptionFilter, Catch, ArgumentsHost } from '@nestjs/common';
import { PixKeyUnknowGrpcError } from 'src/pix-keys/pix-keys.service';
import { Response } from 'express';
import { PixAccountRegisterUnknowGrpcError } from 'src/bank-accounts/bank-accounts.service';
import { PixBankRegisterUnknowGrpcError } from 'src/bank-operations/bank-operations.service';

@Catch(
  PixKeyUnknowGrpcError,
  PixAccountRegisterUnknowGrpcError,
  PixBankRegisterUnknowGrpcError,
)
export class GrpcUnknowErrorFilter implements ExceptionFilter {
  catch(exception: Error, host: ArgumentsHost) {
    const ctx = host.switchToHttp();
    const response = ctx.getResponse<Response>();

    response.status(500).json({
      statusCode: 500,
      message: exception.message,
    });
  }
}
