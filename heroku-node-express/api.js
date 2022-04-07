const express = require("express");
const { Client } = require("pg");

const devDB = "postgres:///webstack_dev";

// env
const conn = process.env.DATABASE_URL || devDB;
const port = process.env.PORT || 3000;

// db
const client = new Client({
  connectionString: conn,
  ssl: conn !== devDB,
});
client.connect();

// routes
const app = express();

app.get("/", (_, resp) => {
  const status = "ok";

  client.query("SELECT 1", (err, _) => {
    if (err !== null) {
      console.log(err);
    }
  });

  resp.json({ status: status });
});

// listen
app.listen(port, () => {
  console.log(`Listening at http://localhost:${port}`);
});
