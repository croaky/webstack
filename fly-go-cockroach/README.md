# Fly Go Cockroach

Develop:

```bash
go run api.go
```

Create a Postgres database:

* Go to <https://cockroachlabs.cloud/> to create a new Postgres database.
* Choose a region to be close to app server.
* Copy "Connection string".

Set up Fly:

```bash
brew install flyctl
fly auth login
rm -rf ~/.docker
fly launch --remote-only
```
