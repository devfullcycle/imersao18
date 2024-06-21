import { Title } from "../components/Title";
import { EventModel } from "../models";
import { EventCard } from "../components/EventCard";

export async function getEvents(): Promise<EventModel[]> {
  const response = await fetch(`${process.env.GOLANG_API_URL}/events`, {
    headers: {
      "apikey": process.env.GOLANG_API_TOKEN as string
    },
    cache: "no-store",
    // next: {
    //   tags: ["events"],
    // }
  });

  return (await response.json()).events;
}

export default async function HomePage() {
  const events = await getEvents();
  console.log(events);
  return (
    <main className="mt-10 flex flex-col">
      <Title>Eventos disponíveis</Title>
      <div className="mt-8 sm:grid sm:grid-cols-auto-fit-cards flex flex-wrap justify-center gap-x-2 gap-y-4">
        {events.map((event) => (
          <EventCard key={event.id} event={event} />
        ))}
      </div>
    </main>
  );
}
