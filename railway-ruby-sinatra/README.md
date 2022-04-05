# Railway Ruby Sinatra

See example deployed at
<https://webstack-ruby-sinatra.up.railway.app/>.

## Setup

```
bundle
bundle exec ruby app.rb
```

Go to <https://railway.app/new>.
Click "Deploy from Repo".
Select this repo.
Select `main` branch.
Click "Deploy".

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
Set "Root directory" to `/railway-ruby-sinatra`.
Set "Healthcheck Path" to `/`.
