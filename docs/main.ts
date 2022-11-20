// :!clear && deno run --allow-read --allow-env --allow-net --allow-write %
import "https://deno.land/x/dotenv/load.ts";
import AsciiTable, { AsciiAlign } from "https://deno.land/x/ascii_table/mod.ts";

const resp = await fetch(
  "https://api.checklyhq.com/v1/reporting?presetWindow=last7Days",
  {
    method: "GET",
    headers: {
      Accept: "application/json",
      Authorization: `Bearer ${Deno.env.get("CHECKLY_API_KEY")}`,
      "X-Checkly-Account": `${Deno.env.get("CHECKLY_ACCOUNT_ID")}`,
    },
  }
);
const data = await resp.json();
var sorted = [];

for (let i = 0; i < data.length; i++) {
  if (data[i]["deactivated"] === false) {
    sorted.push([
      data[i]["name"],
      data[i]["aggregate"]["successRatio"].toFixed(2),
      data[i]["aggregate"]["avg"],
      data[i]["aggregate"]["p95"],
      "y",
    ]);
  }
}

// add historical data for apps I've shut down to save money
sorted.push(["fly-go-cockroach", (100.0).toFixed(2), 246, 648, "n"]);
sorted.push(["fly-go-crunchy-bridge", (100.0).toFixed(2), 254, 574, "n"]);
sorted.push(["fly-go-neon", (100.0).toFixed(2), 361, 510, "n"]);
sorted.push(["fly-go-planetscale", 99.99, 207, 523, "n"]);
sorted.push(["fly-go-postgres-read-replicas", 99.99, 115, 434, "n"]);
sorted.push(["fly-go-sqlite", 99.99, 125, 392, "n"]);
sorted.push(["fly-go-supabase", 99.99, 214, 546, "n"]);
sorted.push(["heroku-go-postgres", 99.99, 303, 829, "n"]);
sorted.push(["railway-go-postgres", 99.99, 493, 642, "n"]);
sorted.push(["render-go-postgres", 99.99, 494, 622, "n"]);

// Sort by uptime as third most important.
// In reality, it's most important but all these services
// are proving to offer 99.99% uptime in practice.
sorted.sort((a, b) => b[1] - a[1]);

// Sort by p95 response time as second most important.
sorted.sort((a, b) => a[3] - b[3]);

// Sort by avg response time as most important.
sorted.sort((a, b) => a[2] - b[2]);

const table = new AsciiTable("API checks");
table.setHeading("stack", "ok %", "avg ms", "p95 ms", "active");
table.setAlign(1, AsciiAlign.RIGHT);
sorted.forEach((row) => table.addRow(row));

const tmplHTML = await Deno.readTextFile("./docs/template.html");
const genHTML = tmplHTML.replace("REPLACE_DASHBOARD_MAIN", table.toString());
await Deno.mkdir("./public", { recursive: true });
await Deno.writeTextFile("./public/index.html", genHTML);
