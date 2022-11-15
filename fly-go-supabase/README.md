# Fly Go w/ Supabase

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

Choose Virginia for a region.
Don't choose to add a Postgres database or deploy now.

Set up Supabase:

* Create a database at [Supabase](https://supabase.com/).
* Go to "Settings" > "Database" > "Connection String" > "URI" > "Copy".

Add connection string to Fly:

```bash
fly secrets set DATABASE_URL="PASTE"
```
