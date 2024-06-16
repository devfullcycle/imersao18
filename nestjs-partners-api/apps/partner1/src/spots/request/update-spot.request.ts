import { PartialType } from '@nestjs/mapped-types';
import { CreateSpotRequest } from './create-spot.request';

export class UpdateSpotRequest extends PartialType(CreateSpotRequest) {}
