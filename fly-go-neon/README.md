# Fly Go

See example deployed to [Fly.io](https://fly.io)
with a [Neon Postgres](https://console.neon.tech/) database
at <https://webstack-go-neon.fly.dev/>.

## Setup

Develop:

```
go run api.go
```

Set up Fly:

```
brew install flyctl
flyctl auth login
rm -rf ~/.docker
flyctl launch --remote-only
```
