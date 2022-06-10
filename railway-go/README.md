# Railway Go

See example deployed at
<https://webstack-go.up.railway.app/>.

## Setup

```
go run api.go
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
Set "Root directory" to `/railway-go`.
Set "Healthcheck Path" to `/`.
