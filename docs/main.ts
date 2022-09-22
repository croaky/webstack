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
  },
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
    ]);
  }
}

sorted.sort((a, b) => a[2] - b[2]);

const table = new AsciiTable("API checks last 7 days");
table.setHeading("Name", "OK (%)", "Avg (ms)", "p95 (ms)");
table.setAlign(1, AsciiAlign.RIGHT);
sorted.forEach((row) => table.addRow(row));

const tmplHTML = await Deno.readTextFile("./docs/template.html");
const genHTML = tmplHTML.replace("REPLACE_DASHBOARD_MAIN", table.toString());
await Deno.mkdir("./public", { recursive: true });
await Deno.writeTextFile("./public/index.html", genHTML);
