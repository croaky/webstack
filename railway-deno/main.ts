import { serve } from "https://deno.land/std@0.143.0/http/server.ts";

const port = Number(Deno.env.get("PORT") || 8080);

const handler = (_req: Request): Response => {
  return new Response(`{"status":"ok"}`, {
    headers: { "Content-Type": "application/json" },
    status: 200,
  });
};

console.log(`Listening at http://localhost:${port}/`);
await serve(handler, { port });
