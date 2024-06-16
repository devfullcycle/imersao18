import { PartialType } from '@nestjs/mapped-types';
import { CriarLugarRequest } from './criar-lugar.request';

export class AtualizarLugarRequest extends PartialType(CriarLugarRequest) {}
