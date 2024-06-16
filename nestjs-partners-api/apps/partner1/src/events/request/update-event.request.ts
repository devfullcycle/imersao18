import { PartialType } from '@nestjs/mapped-types';
import { CreateEventRequest } from './create-event.request';

export class UpdateEventRequest extends PartialType(CreateEventRequest) {}
