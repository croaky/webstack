# Fly Crunchy Bridge Read Replicas

This stack involves:

- a web application written in Go
- deployed to 2 regions on Fly.io
- querying a Postgres primary and read replica on Crunchy Bridge
- networked with Tailscale

Develop:

```bash
DATABASE_URL="postgres:///webstack_dev" go run api.go
```

Set up Crunchy Bridge:

- Go to [Crunchy Bridge](https://crunchybridge.com/).
- Create a cluster in AWS `us-west-1` (N. California).
- Click "Connection" tab.
- Select "Role: postgres" and "Format: psql".
- Click "Copy", then paste into a shell:

```bash
psql postgres://postgres:password@host.db.postgresbridge.com:5432/postgres
```

In the Postgres shell as the `postgres` superuser,
[enable PgBouncer](https://docs.crunchybridge.com/how-to/pgbouncer/):

```sql
CREATE EXTENSION crunchy_pooler;
```

Create a reusable, ephemeral
[Tailscale auth key](https://login.tailscale.com/admin/settings/keys).

In the Crunchy Bridge web UI:

- Click "Networking".
- Delete the two "fully open firewall" rules `::/0` and `0.0.0.0/0`.
- Click "Tailscale".
- Paste the Tailscale auth key.S
- Click "Connect Tailscale".

In the Tailscale web UI:

- Go to the [Machines](https://login.tailscale.com/admin/machines) page.
- Click the three dots next to the newly-connected Crunchy database.
- Click "Edit machine name..".
- Rename the machine something like `crunchy-n-california`.
- Click the machine name.
- Copy the domain name into a temporary text file.
  It will look something like `crunchy-n-california.taile1234.ts.net`.

In the Crunchy Bridge web UI:

- Click "Connection"
- Select "Role: application".
- Copy the connection string and paste into the temporary text file.
- Replace the Crunchy Bridge host (e.g. `p.abc123.db.postgresbridge.com`)
  with the Tailscale domain name copied above.
- Replace port `5432` with `5431` so the connection will use PgBouncer.
- Copy the edited connection string, which will look something like:

```
postgres://application:password@crunchy-n-california.taile1234.ts.net:5431/postgres"
```

Create Fly app:

```bash
fly launch --machine remote-only
```

Set the primary database URL connection string:

```bash
fly secrets set DATABASE_URL="<paste>"
```

Add a London region near the read replica:

```bash
fly regions add lhr
```

In the Crunchy Bridge web UI from the cluster overview page:

- Click "Actions".
- Click "Create Replica".
- Select AWS region "EU West 2 (London)".
- Set an environment variable on Fly for the London region,
  using the Crunchy `application` user and `5431` PgBouncer port:

```bash
fly secrets set LONDON_URL="<london-connection-string>"
```

Deploy:

```bash
fly deploy --remote-only
```

Scale 1 instance per region:

```bash
fly autoscale set min=2 max=2
```
