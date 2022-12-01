# Railway Go Crunchy Bridge

This stack involves:

- a web application written in Go
- deployed to Railway
- querying a Postgres database on Crunchy Bridge
- networked with Tailscale

Develop:

```bash
createdb webstack_dev
DATABASE_URL="postgres:///webstack_dev" go run api.go
```

Set up Crunchy Bridge:

- Go to [Crunchy Bridge](https://crunchybridge.com/).
- Create a cluster in AWS `us-east-1` (N. Virginia).
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
- Paste the Tailscale auth key.
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
postgres://application:password@crunchy-n-california.taile1234.ts.net:5431/postgres
```

Go to <https://railway.app/new>:

- Click "Deploy from Repo".
- Select this repo.
- Select `main` branch.
- Click "Deploy now".
- Click "Settings" tab.
- Set "Root directory" to `/railway-go-crunchy-bridge`.
- Click "Variables".
- Click "+ New Variable".
- Enter "DATABASE_URL" and paste in the Tailscale'd Crunchy connection string.
