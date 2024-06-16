import { PartialType } from '@nestjs/mapped-types';
import { CriarEventoRequest } from './criar-evento.request';

export class AtualizarEventoRequest extends PartialType(CriarEventoRequest) {}
