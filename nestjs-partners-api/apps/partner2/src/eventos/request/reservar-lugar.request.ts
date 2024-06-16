export class ReservarLugarRequest {
  lugares: string[]; //['A1', 'A2']
  tipo_ingresso: 'inteira' | 'meia';
  email: string;
}
