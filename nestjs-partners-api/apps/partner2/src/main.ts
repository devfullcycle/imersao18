import { NestFactory } from '@nestjs/core';
import { Partner2Module } from './partner2.module';

async function bootstrap() {
  const app = await NestFactory.create(Partner2Module);
  await app.listen(3001);
}
bootstrap();
