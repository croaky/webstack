# Heroku Go Gin

## Setup

```
go get -u github.com/gin-gonic/gin
go run api.go
```

Go to <https://dashboard.heroku.com/apps>.
Click "New" > "Create new app".
Give it a name and submit.

```
heroku login
heroku git:remote -a <app>
heroku buildpacks:add -a <app> https://github.com/lstoll/heroku-buildpack-monorepo
heroku buildpacks:add -a <app> heroku/go
heroku config:set APP_BASE=heroku-go-gin -a webstack-go-gin
```
