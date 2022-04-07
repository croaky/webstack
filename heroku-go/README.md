# Heroku Go

See example deployed at
<https://webstack-go.herokuapp.com/>.

## Setup

```
go run api.go
```

Go to <https://dashboard.heroku.com/apps>.
Click "New" > "Create new app".
Give it a name and submit.
Go to "Deploy" tab and configure GitHub deploys.

```
heroku login
heroku git:remote -a <app>
git remote rename heroku heroku-go
heroku buildpacks:add -a <app> https://github.com/lstoll/heroku-buildpack-monorepo
heroku buildpacks:add -a <app> heroku/go
heroku config:set APP_BASE=heroku-go -a <app>
heroku addons:create heroku-postgresql:hobby-dev -a <app>
```

## Editorial

* I always forget to put the magic comment `// +heroku goVersion 1.18`
  at the top of `go.mod`. Why isn't `go 1.18` later in the file enough?
* From `git push` to build finishing and process restarting is ~10s.
* Databases are backed up continuously.
* HA databases available.
