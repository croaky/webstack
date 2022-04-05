# Railway Ruby X

See example deployed at
<https://webstack-ruby-x.up.railway.app/>.

## Setup

Develop locally:

```
bundle
createdb webstack_dev
bundle exec ruby api.rb
```

Setup Railway:

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
Set "Root directory" to `/railway-ruby-x`.
Set "Healthcheck Path" to `/`.

## Editorial

* The web UI is outstanding.
* Failed deploys were easy to debug.
* The first successful deployment took about.
