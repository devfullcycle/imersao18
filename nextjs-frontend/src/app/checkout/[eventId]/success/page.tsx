import { cookies } from "next/headers";
import { Title } from "../../../../components/Title";
import { EventModel } from "../../../../models";
// queries
export async function getEvent(eventId: string): Promise<EventModel> {
  const response = await fetch(`${process.env.GOLANG_API_URL}/events/${eventId}`, {
    headers: {
      "apikey": process.env.GOLANG_API_TOKEN as string
    },
    cache: "no-store",
    next: {
      tags: [`events/${eventId}`],
    }
  });

  return response.json();
}

export default async function CheckoutSuccessPage({
  params,
}: {
  params: { eventId: string };
}) {
  const event = await getEvent(params.eventId);
  const cookiesStore = cookies();
  const selectedSpots = JSON.parse(cookiesStore.get("spots")?.value || "[]");
  return (
    <main className="mt-10 flex flex-col flex-wrap items-center ">
      <Title>Compra realizada com sucesso!</Title>
      <div className="mb-4 flex max-h-[250px] w-full max-w-[478px] flex-col gap-y-6 rounded-2xl bg-secondary p-4">
        <Title>Resumo da compra</Title>
        <p className="font-semibold">
          Evento {event.name}
          <br />
          Local {event.location}
          <br />
          Data{" "}
          {new Date(event.date).toLocaleDateString("pt-BR", {
            weekday: "long",
            day: "2-digit",
            month: "2-digit",
            year: "numeric",
          })}
        </p>
        <p className="font-semibold text-white">Lugares escolhidos: {selectedSpots.join(", ")}</p>
        
      </div>
    </main>
  );
}
