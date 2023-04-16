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

Set up Neon:

- [Create a project](https://console.neon.tech/app/projects).
- Copy the `psql` string value.
- Append `sslmode=verify-full`

Set the primary database URL connection string:

```bash
fly secrets set DATABASE_URL="<paste>"
```

Deploy:

```bash
fly deploy --remote-only
```
