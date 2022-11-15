# Fly Read Replicas

Develop:

```bash
bundle
createdb webstack_dev
DATABASE_URL="postgres:///webstack_dev" bundle exec ruby api.rb
```

Set up Fly:

```bash
brew install flyctl
flyctl auth login
```

I prefer to not install Docker on my local machine.
One "gotcha" is a leftover `~/.docker` directory can cause errors on Fly.
Remove it:

```bash
rm -rf ~/.docker
```

Now we can use Fly.io's `--remote-only` option.
It will use a remote Docker builder that they set up in our account:

```bash
flyctl launch --remote-only
```

Configure primary region where the primary database is:

```bash
flyctl secrets set PRIMARY_REGION=sjc
```

Create the regions where the read replicas will go:

```bash
fly regions add ams lhr syd yul
```

Create [read replicas](https://fly.io/docs/getting-started/multi-region-databases/#connect-to-regional-replicas)
in those regions:

```bash
fly volumes create pg_data -a webstack-read-replicas-db --size 1 --region ams
fly volumes create pg_data -a webstack-read-replicas-db --size 1 --region lhr
fly volumes create pg_data -a webstack-read-replicas-db --size 1 --region syd
fly volumes create pg_data -a webstack-read-replicas-db --size 1 --region yul
```

Scale 1 instance per region:

```bash
fly autoscale standard min=5 max=5
```
