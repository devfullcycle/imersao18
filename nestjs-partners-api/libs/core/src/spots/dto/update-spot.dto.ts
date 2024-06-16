import { PartialType } from '@nestjs/mapped-types';
import { CreateSpotDto } from './create-spot.dto';

export class UpdateSpotDto extends PartialType(CreateSpotDto) {}
