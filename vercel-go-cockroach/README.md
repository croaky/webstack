# Vercel Go w/ Cockroach serverless Postgres

Vercel supports [serverless functions written in
Go](https://vercel.com/docs/concepts/functions/supported-languages#go).

## Setup

* Go to <https://vercel.com/dashboard>.
* Click "New project".
* Select this Git repo.
* Select `vercel-go` as root directory.
* Click "Deploy".
* Test your app in a web browser at the `/api` path.

Create a Postgres database:

* Go to <https://cockroachlabs.cloud/> to create a new Postgres database.
* Choose "AWS > N. Virginia" region
  to be close to Vercel Washington DC functions.
* Copy "Connection string".

Back in Vercel:

* Click "Settings" tab.
* Click "Environment Variables"
* Create a new env var `DATABASE_URL` and paste value from Neon.

## Editorial

Vercel:

* Preview deploys in branches and pull requests are nice.
* Deploys take ~20s.
* Build and deployment logs are easy to read.
* My functions appear to be in Washington, DC:

```
Generated build outputs:
 - Static files: 3
 - Serverless Functions: 1
 - Edge Functions: 0
Serverless regions: Washington, D.C., USA
```

Cockroach Serverless:

* Impressive low latency. Lowest I've been able to achieve so far for any
  serverless function w/ SQL database.
* Unclear under what conditions it idles.
* `sslmode=verify-full` works and is the default in Go template.
* Not totally compatible w/ Postgres,
  doesn't support Postgres functions or triggers.
