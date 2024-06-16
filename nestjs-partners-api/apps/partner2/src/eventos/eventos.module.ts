import { Module } from '@nestjs/common';
import { EventsCoreModule } from '@app/core/events/events-core.module';
import { EventosControllers } from './eventos.controller';

@Module({
  imports: [EventsCoreModule],
  controllers: [EventosControllers],
})
export class EventosModule {}
