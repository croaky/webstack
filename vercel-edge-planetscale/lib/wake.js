import fetch from "node-fetch";
import { Client } from "@planetscale/database";
import * as dotenv from "dotenv";

dotenv.config();

const db = new Client({
  fetch: fetch,
  url: process.env["DATABASE_URL"],
});

(async () => {
  const conn = db.connection();
  await conn.execute("INSERT INTO checks (at) VALUES (now())");
})();
