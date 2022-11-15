# Railway Ruby

Setup:

```bash
bundle
createdb webstack_dev
DATABASE_URL="postgres:///webstack_dev" bundle exec ruby api.rb
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
Set "Root directory" to `/railway-ruby`.
Set "Healthcheck Path" to `/`.
