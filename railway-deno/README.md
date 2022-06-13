# Railway Deno

See example deployed at
<https://webstack-deno.up.railway.app/>.

## Setup

Create a new Railway project with a web service:

* Go to your [Railway dashboard](https://railway.app/dashboard).
* Click "+ New Project".
* Select "Deploy from GitHub repo".
* Select this repo.
* Select "Deploy Now". The first build will fail.
* Click "Settings".
* Click "Generate Domain" to expose the service to the public internet.
* Fill in "Root Directory" with `/railway-deno`.
* Fill in "Start Command" with `deno run --allow-env --allow-net ./railway-deno/main.ts`.
* Leave "Healthcheck Path" blank for now.
* For "Builder", select "Nixpacks".
* Type "Esc".

Add a Postgres database to the Railway project:

* Click "+ New"
* Click "Database"
* Click "Add PostgreSQL"
* Click the "PostgreSQL" rectangle on the Railway palette.
* Click "Connect".
* Copy "Postgres Connection URL".
* Type "Esc".
* Click your web service rectangle on the Railway palette.
* Click "Variables".
* Enter `DATABASE_URL` for the variable name and paste the value.
* Append `?sslmode=disable` to the value to avoid the error noted below.
* Click "Add".

## Editorial

The Deno Postgres database driver threw this error when trying to connect
to the Railway Postgres instance:

```
Sending fatal alert BadCertificate
TLS connection failed with message: invalid peer certificate contents: invalid peer certificate: UnknownIssuer
Defaulting to non-encrypted connection
```

I didn't see a similar issue for Go's or Ruby's Postgres drivers on Railway.
Obviously, it would be preferable to not disable SSL!
