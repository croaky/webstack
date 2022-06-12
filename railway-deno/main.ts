import { serve } from "https://deno.land/std@0.143.0/http/server.ts";
import { Pool } from "https://deno.land/x/postgres@v0.16.1/mod.ts";

// env
const port = Number(Deno.env.get("PORT") || 8080);
const dbUrl =
  Deno.env.get("DATABASE_URL") ||
  "postgres://postgres:postgres@0.0.0.0:5432/webstack_dev";

// db
const pool = new Pool(dbUrl, 5, true);
const db = await pool.connect();

// routes
const handler = async (_req: Request): Promise<Response> => {
  await db.queryArray("SELECT 1");

  return new Response(`{"status":"ok"}`, {
    headers: { "Content-Type": "application/json" },
    status: 200,
  });
};

// listen
console.log(`Listening at http://localhost:${port}/`);
await serve(handler, { port });
