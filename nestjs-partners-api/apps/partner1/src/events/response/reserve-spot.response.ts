import { Spot, Ticket } from '@prisma/client';

type TicketWithSpots = Ticket & { Spot: Spot };

export class ReserveSpotResponse {
  constructor(readonly tickets: TicketWithSpots[]) {}

  toJSON() {
    return this.tickets.map((ticket) => ({
      id: ticket.id,
      email: ticket.email,
      spot: ticket.Spot.name,
      ticket_kind: ticket.ticketKind,
      status: ticket.Spot.status,
      event_id: ticket.Spot.eventId,
    }));
  }
}
