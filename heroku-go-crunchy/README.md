# Heroku Crunchy

This stack involves:

- a web application written in Go
- deployed to Heroku
- connecting to a Postgres database via PgBouncer on Crunchy Bridge
- networked with Tailscale

Develop:

```bash
createddb webstack_dev
DATABASE_URL="postgres:///webstack_dev" go run api.go
```

Set up Crunchy Bridge:

- Go to [Crunchy Bridge](https://crunchybridge.com/).
- Create a cluster in AWS `us-east-1` (Virginia).
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
- Rename the machine something like `crunchy-virginia`.
- Click the machine name.
- Copy the domain name into a temporary text file.
  It will look something like `crunchy-virginia.taile1234.ts.net`.

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

Go to <https://dashboard.heroku.com/apps>:

- Click "New" > "Create new app".
- Give it a name and submit.
- Click "GitHub" and select this repo, then "Connect".
- Click "Enable Automatic Deploys"

```bash
heroku login
heroku git:remote -a <app>
git remote rename heroku heroku-go-crunchy
heroku config:set TAILSCALE_AUTHKEY=<paste> -a <app>
heroku labs:enable runtime-dyno-metadata -a <app>
heroku stack:set container -a <app>
```

Commit to the repo to deploy.
