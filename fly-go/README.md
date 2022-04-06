# Fly Go

## Setup

Install Docker from <https://www.docker.com/get-started/> (~1.5GB for M1 chip).

```
flyctl launch
```

## Editorial

I tried using the service without having Docker installed locally and was unable
to deploy successfully. Example errors below.

Error on `flyctl launch`:

```
Error Failed attaching webstack-go-db to the Postgres cluster webstack-go: unexpected end of JSON input.\nTry attaching manually with 'fly postgres attach --app webstack-go --postgres-app webstack-go-db'
```

Error on `flyctl deploy`:

```
v0 failed - Failed due to unhealthy allocations - no stable job version to auto revert to and deploying as v1
```

After installing Docker, `flyctl launch` worked smoothly.
