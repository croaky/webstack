import { Client } from "@planetscale/database";

export const config = {
  runtime: "experimental-edge",
};

const db = new Client({
  url: process.env["DATABASE_URL"],
});

export default async function handler(_req: Request) {
  const conn = db.connection();
  await conn.execute("INSERT INTO checks (at) VALUES (now())");

  return new Response('{"status":"ok"}', {
    headers: {
      "Content-Type": "application/json",
    },
  });
}
