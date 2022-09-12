# Fly Go

See example deployed at
<https://webstack-go.fly.dev/>.

## Setup

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

I prefer to not install Docker on my local machine. Fly.io has a cool
`--remote-only` option that will use a remote Docker builder that they set up in
your account. One "gotcha" is a leftover `~/.docker` directory can cause these
errors when using Fly:

```
Error failed to fetch builder image 'index.docker.io/heroku/buildpacks:20': resolve auth for ref index.docker.io/heroku/buildpacks:20: error getting credentials - err: exec: "docker-credential-desktop": executable file not found in $PATH, out: ``
```
