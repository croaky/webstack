# Railway Ruby Sinatra Postgres

See example deployed at
<https://webstack-production.up.railway.app/>.

## Setup

Go to <https://railway.app/new>.
Click "Deploy from Repo".
Select this repo.
Select `main` branch.
Click "Deploy".

It will fail.

Click "+ New" button.
Click "Database".
Click "Add PostgreSQL".
Click "Connect" tab.
Copy "Postgres Connection URL".

Click the original web service.
Click "Variables" tab.
Click "+ New Variable".
Enter `DATABASE_URL` and paste copied string.

Click "Settings" tab.
Set "Root directory" to `/railway-ruby-sinatra-pg`.
Set "Healthcheck Path" to `/`.

## Editorial

* The web UI is outstanding.
* Failed deploys were easy to debug.
* The first successful deployment took about.
