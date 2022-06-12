# Railway Deno

See example deployed at
<https://webstack-deno.up.railway.app/>.

* Go to your [Railway dashboard](https://railway.app/dashboard).
* Click "+ New Project".
* Select "Deploy from GitHub repo".
* Select this repo.
* Select "Deploy Now". The first build will fail.
* Click "Settings".
* Fill in "Root Directory" with `/railway-deno`.
* Fill in "Start Command" with `deno run --allow-env --allow-net ./railway-deno/main.ts`.
* Leave "Healthcheck Path" blank for now.
* For "Builder", select "Nixpacks".
* Click "Deployments" to watch a deploy succeed.

Once the deploy is successful:

* Click "Settings".
* Click "Generate Domain" to expose the service to the public internet.
