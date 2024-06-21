import { cookies } from "next/headers";
import { Title } from "../../components/Title";
import { redirect } from "next/navigation";
import { EventModel } from "../../models";
import { CheckoutForm } from "./CheckoutForm";

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

export default async function CheckoutPage() {
  const cookiesStore = cookies();
  const eventId = cookiesStore.get("eventId")?.value;
  if (!eventId) {
    return redirect("/");
  }
  const event = await getEvent(eventId);
  const selectedSpots = JSON.parse(cookiesStore.get("spots")?.value || "[]");
  let totalPrice = selectedSpots.length * event.price;
  const ticketKind = cookiesStore.get("ticketKind")?.value;
  if (ticketKind === "half") {
    totalPrice = totalPrice / 2;
  }
  const formattedTotalPrice = new Intl.NumberFormat("pt-BR", {
    style: "currency",
    currency: "BRL",
  }).format(totalPrice);
  return (
    <main className="mt-10 flex flex-wrap justify-center md:justify-between">
      <div className="mb-4 flex max-h-[250px] w-full max-w-[478px] flex-col gap-y-6 rounded-2xl bg-secondary p-4">
        <Title>Resumo da compra</Title>
        <p className="font-semibold">
          {event.name}
          <br />
          {event.location}
          <br />
          {new Date(event.date).toLocaleDateString("pt-BR", {
            weekday: "long",
            day: "2-digit",
            month: "2-digit",
            year: "numeric",
          })}
        </p>
        <p className="font-semibold text-white">{formattedTotalPrice}</p>
      </div>
      <div className="w-full max-w-[650px] rounded-2xl bg-secondary p-4">
        <Title>Informações de pagamento</Title>
        <CheckoutForm className="mt-6 flex flex-col gap-y-3">
          <div className="flex flex-col">
            <label htmlFor="titular">E-mail</label>
            <input
              type="email"
              name="email"
              className="mt-2 border-solid rounded p-2 h-10 bg-input"
              defaultValue={"test@test.com"}
            />
          </div>
          <div className="flex flex-col">
            <label htmlFor="titular">Nome no cartão</label>
            <input
              type="text"
              name="card_name"
              className="mt-2 border-solid rounded p-2 h-10 bg-input"
              defaultValue={"Teste Teste"}
            />
          </div>
          <div className="flex flex-col">
            <label htmlFor="cc">Numero do cartão</label>
            <input
              type="card_number"
              name="cc"
              className="mt-2 border-solid rounded p-2 h-10 bg-input"
              defaultValue={"4111111111111111"}
            />
          </div>
          <div className="flex flex-wrap sm:justify-between">
            <div className="flex w-full flex-col md:w-auto">
              <label htmlFor="expire">Vencimento</label>
              <input
                type="text"
                name="expire_date"
                className="mt-2 sm:w-[240px] border-solid rounded p-2 h-10 bg-input"
                defaultValue={"12/2024"}
              />
            </div>
            <div className="flex w-full flex-col md:w-auto">
              <label htmlFor="cvv">CVV</label>
              <input
                type="text"
                name="cvv"
                className="mt-2 sm:w-[240px] border-solid rounded p-2 h-10 bg-input"
                defaultValue={"123"}
              />
            </div>
          </div>
          <button className="rounded-lg bg-btn-primary py-4 px-4 text-sm font-semibold uppercase text-btn-primary">
            Finalizar pagamento
          </button>
        </CheckoutForm>
      </div>
    </main>
  );
}
