import { Module } from '@nestjs/common';
import { SpotsService } from './spots.service';

@Module({
  providers: [SpotsService],
  exports: [SpotsService],
})
export class SpotsCoreModule {}
