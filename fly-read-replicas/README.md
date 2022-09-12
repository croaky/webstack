# Fly Ruby

See example deployed at
<https://webstack-read-replicas.fly.dev/>.

## Setup

Develop:

```
bundle
createdb webstack_dev
DATABASE_URL="postgres:///webstack_dev" bundle exec ruby api.rb
```

Set up Fly:

```
brew install flyctl
flyctl auth login
rm -rf ~/.docker
flyctl launch --remote-only
flyctl secrets set PRIMARY_REGION=sjc
```

I prefer to not install Docker on my local machine. Fly.io has a cool
`--remote-only` option that will use a remote Docker builder that they set up in
your account. One "gotcha" is a leftover `~/.docker` directory can cause these
errors when using Fly:

```
Error failed to fetch builder image 'index.docker.io/heroku/buildpacks:20': resolve auth for ref index.docker.io/heroku/buildpacks:20: error getting credentials - err: exec: "docker-credential-desktop": executable file not found in $PATH, out: ``
```
