# Fly Go PlanetScale

See example deployed at
<https://webstack-go-planetscale.fly.dev/>.

## Setup

Develop:

```
go run api.go
```

Set up Fly:

```
brew install flyctl
fly auth login
rm -rf ~/.docker
fly launch --remote-only
```
