# Vercel Go

See example deployed at
<https://webstack-go.vercel.app/api>.

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

* Go to <https://neon.tech> to create a new Postgres database.
* Enable connection pooling.
* Copy "Connection string".

Back in Vercel:

* Click "Settings" tab.
* Click "Environment Variables"
* Create a new env var `DATABASE_URL` and paste value from Neon.

## Editorial

* Preview deploys in branches and pull requests are nice.
* Go 1.18 was released 2022-03-15. As of 2022-04-04, it is not supported.
* Deploys take ~20s.
* Deployment logs are easy to read.
