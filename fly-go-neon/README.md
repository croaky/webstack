# Fly Neon

[Neon Postgres](https://console.neon.tech/).

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
