import { serve } from "https://deno.land/std@0.114.0/http/server.ts";
import * as postgres from "https://deno.land/x/postgres@v0.15.0/mod.ts";

const dbURL = Deno.env.get("DATABASE_URL")!;
const pool = new postgres.Pool(dbURL, 3, true);

serve(async (req) => {
  const db = await pool.connect();

  try {
    await db.queryArray`SELECT 1`;
    return new Response('{"status": "ok"}', {
      headers: { "content-type": "application/json" },
    });
  } catch (err) {
    console.error(err);
    return new Response(`Internal Server Error\n\n${err.message}`, {
      status: 500,
    });
  } finally {
    db.release();
  }
});
