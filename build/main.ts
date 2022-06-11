import "https://deno.land/x/dotenv/load.ts";
import AsciiTable, { AsciiAlign } from "https://deno.land/x/ascii_table/mod.ts";

const resp = await fetch(
  "https://api.checklyhq.com/v1/reporting?presetWindow=last7Days&deactivated=false",
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

const table = new AsciiTable("Web stacks last 7 days");
table.setHeading("Name", "OK %", "Avg (ms)", "p95 (ms)");
table.setAlign(1, AsciiAlign.RIGHT);

for (let i = 0; i < data.length; i++) {
  table.addRow(
    data[i]["name"],
    `${data[i]["aggregate"]["successRatio"].toFixed(1)}%`,
    data[i]["aggregate"]["avg"],
    data[i]["aggregate"]["p95"]
  );
}

console.log(table.toString());
