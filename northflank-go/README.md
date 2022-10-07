# Northflank Go

See example deployed at
<https://site--us-central--webstack-go--rluc-sxvr.code.run/>.

## Setup

Develop:

```
go run api.go
```

Create Northflank project:

* [Create a project](https://northflank.com/docs/v1/application/getting-started/create-a-project)
* Choose a name and region "US - Central" or "Europe - West".

Add Postgres database:

* Click "Addons" > "Create new addon" > "PostgreSQL".
* Give it a name, choose Postgres v 14.5.0.
* Keep networking defaults: "Deploy with TLS" enabled and "Publicly accessible" disabled.
* Click "Create addon".

Expose Postgres connection to project via secret group:

* Click "Connection details" > "Link to secret groups" > "Create a new secret group".
* Give it a name.
* Within "Linked addons", click "Show addons".
* Click "Configure" next to the Postgres database.
* Click `POSTGRES_URI` > "Aliases" and enter `DATABASE_URL`.
* Click "Create secret group".

Create new service:

* Choose "Add new service".
* Choose "Build: Build a Git repo and store image".
* Choose GitHub repository and `main` branch.
* For build options, choose "Buildpack" > `heroku/buildpacks:20` and "Build context" of `/northflank-go`.
* In "Advanced", "CMD override" to `go run api.go`.
* Click "Create service".

## Editorial

Little things I came across:

* I could get <code>heroku/buildpacks:20</code> working, but not
  <code>heroku/builder:22</code>.
* HTTP/2 works for the app from browser to Northflank's routing layer.
  I had to use the default HTTP/1.1 option in Northflank, which confused me
  but is specific to Go's HTTP package and what you'd need to do to get TLS
  working from Northflank's router to your app process.
* Within "Services" > "Builds", when I click "Start build",
  I'm redirected to the "Commits" tab and no build seems to start.
* I'm confused about how I set up pull request deploys and staging-production
  pipelines. "Combined" services don't let me select `*` pull requests but let
  me auto-deploy `main`. In both "Build" and "Combined" services,
  I haven't figured out yet how to add a service to a pipeline. I can select
  select a service in a "Build" pipeline but am unable to add.
