import { Spot, Ticket } from '@prisma/client';

type TicketWithSpots = Ticket & { Spot: Spot };

export class ReservarLugarResponse {
  constructor(readonly tickets: TicketWithSpots[]) {}

  toJSON() {
    return this.tickets.map((ticket) => ({
      id: ticket.id,
      email: ticket.email,
      lugar: ticket.Spot.name,
      tipo_ingresso: ticket.ticketKind,
      estado: ticket.Spot.status,
      evento_id: ticket.Spot.eventId,
    }));
  }
}
