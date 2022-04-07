# Render Go

See example deployed at
<https://webstack-go.onrender.com/>.

## Setup

```
go run api.go
```

Go to <https://dashboard.render.com/new/database>
Name the database "webstack".
Set "Plans" to "Free".
Copy "Internal Connection String".

Go to <https://dashboard.render.com/select-repo?type=web>.
Select this repo.
Select `main` branch.
Name the service.
Select "Go" environment.
Set "Build command" to `cd render-go && go build -tags netgo -ldflags '-s -w' -o api`.
Set "Start command" to `cd render-go && ./api`.
Set "Plans" to "Free".
Click "Advanced".
Click "Add Environment Variable".
Paste "Internal Connection String" value from above into new `DATABASE_URL` env var.

