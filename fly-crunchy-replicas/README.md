# Fly Crunchy Bridge Read Replicas

Develop:

```bash
DATABASE_URL="postgres:///webstack_dev" go run api.go
```

Set up Crunchy Bridge:

* Go to [Crunchy Bridge](https://crunchybridge.com/).
* Create a cluster in AWS `us-west-1` (N. California).
* Click "Connection" tab.
* Select "Role: postgres" and "Format: psql".
* Click "Copy", then paste into a shell:

```bash
psql postgres://postgres:password@host.db.postgresbridge.com:5432/postgres
```

In the Postgres shell,
[enable PgBouncer](https://docs.crunchybridge.com/how-to/pgbouncer/):

```sql
CREATE EXTENSION crunchy_pooler;
```

Back in the Crunchy Bridge web UI,
select "Role: application",
copy the connection string again,
and replace port `5432` with `5431`.

Create Fly app:

```bash
fly launch --remote-only
```

Set the primary database URL connection string:

```bash
fly secrets set DATABASE_URL="<paste>"
```

Create the regions near the read replicas:

```bash
fly regions add ams syd yul
```

Create read replicas near each region and
set an environment variable on Fly for each one,
using the Crunchy `application` user and `5431` PgBouncer port:

* AWS eu-north-1 (Stockholm)
* AWS ap-southeast-2 (Sydney)
* AWS ca-central-1 (Toronto)

```bash
fly secrets set EUROPE_URL="<stockholm-connection-string>"
fly secrets set AUSTRALIA_URL="<sydney-connection-string>"
fly secrets set CANADA_URL="<toronto-connection-string>"
```

Scale 1 instance per region:

```bash
fly autoscale set min=4 max=4
```

Deploy:

```bash
fly deploy --remote-only
```
