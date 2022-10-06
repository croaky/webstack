# Vercel Go w/ Neon Postgres database

See example deployed at
<https://webstack-go-neon.vercel.app/api>.

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
* Ensure connection pooling is disabled
* Copy basic `psql` connection string.

Back in Vercel:

* Click "Settings" tab.
* Click "Environment Variables"
* Create a new env var `DATABASE_URL` and paste value from Neon.
  Append `?sslmode=verify-full&options=project%3DYOUR-PROJECT-ID-HERE`.
  See [project ID in options docs](https://neon.tech/docs/how-to-guides/connectivity-issues/#a-pass-project-id-in-options)

## Editorial

Vercel:

* Preview deploys in branches and pull requests are nice.
* Deploys take ~20s.
* Build and deployment logs are easy to read.

Neon:

* Pretty awesome that the database can "scale to zero",
  idles after 5m of inactivity.
* `sslmode=verify-full` works.
* Normal Postgres, uses version 14.5 as of Oct 6, 2022.
