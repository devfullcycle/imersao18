import { revalidateTag } from "next/cache";
import { NextRequest, NextResponse } from "next/server";

export function POST(request: NextRequest, {params}: {params: {eventId: string}}) {
  
    revalidateTag("events");
    revalidateTag(`events/${params.eventId}`);

    return new Response(null, {status: 204});
}
