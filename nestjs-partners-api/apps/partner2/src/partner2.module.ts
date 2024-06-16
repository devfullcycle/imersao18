import { Module } from '@nestjs/common';
import { ConfigModule } from '@nestjs/config';
import { EventosModule } from './eventos/eventos.module';
import { PrismaModule } from '@app/core/prisma/prisma.module';
import { LugaresModule } from './lugares/lugares.module';
import { AuthModule } from '../../../libs/core/src/auth/auth.module';

@Module({
  imports: [
    ConfigModule.forRoot({
      envFilePath: '.env.partner2',
      isGlobal: true,
    }),
    AuthModule,
    PrismaModule,
    EventosModule,
    LugaresModule,
  ],
})
export class Partner2Module {}
