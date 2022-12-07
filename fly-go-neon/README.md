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
fly auth login
fly launch --remote-only
```
