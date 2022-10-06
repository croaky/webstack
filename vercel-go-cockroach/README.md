# Vercel Go w/ Cockroach serverless Postgres

See example deployed at
<https://webstack-go-cockroach.vercel.app/api>.

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
* Copy "Connection string".

Back in Vercel:

* Click "Settings" tab.
* Click "Environment Variables"
* Create a new env var `DATABASE_URL` and paste value from Neon.
