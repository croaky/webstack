# Render Go

Develop:

```bash
go run main.go
```

Postgres:

* Go to <https://dashboard.render.com/new/database>
* Create a new database.
* Select a "Plan".
* Once the database is available, copy "Internal Connection String".

Web service:

* Go to <https://dashboard.render.com/select-repo?type=web>.
* Select this repo.
* Select `main` branch.
* Name the service.
* Select "Go" environment.
* Set "Build command" to `cd render-go && go build -tags netgo -ldflags '-s -w' -o api`.
* Set "Start command" to `cd render-go && ./api`.
* Set "Plans" to "Free".
* Click "Advanced".
* Click "Add Environment Variable".
* Paste "Internal Connection String" value from Postgres section above into new `DATABASE_URL` env var
