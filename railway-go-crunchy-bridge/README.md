# Railway Go Crunchy Bridge

This stack involves:

- a web application written in Go
- deployed to Railway
- querying a Postgres database on Crunchy Bridge
- networked with Tailscale

Develop:

```bash
go run api.go
```

Create a Postgres database:

- Go to <https://crunchybridge.com/> to create a new Postgres database.
- Choose a region to be close to app server.
- Copy "Connection string".

Go to <https://railway.app/new>:

- Click "Deploy from Repo".
- Select this repo.
- Select `main` branch.
- Click "Deploy".
- Click "Settings" tab.
- Set "Root directory" to `/railway-go-crunchy-bridge`.
- Set "Healthcheck Path" to `/`.
