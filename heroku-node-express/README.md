# Heroku Node Express

## Setup

```
npm install
createdb webstack_dev
DATABASE_URL="postgres:///webstack_dev" node api.js
```

Go to <https://dashboard.heroku.com/apps>.
Click "New" > "Create new app".
Give it a name and submit.
Click "GitHub" and select this repo, then "Connect".
Click "Enable Automatic Deploys"

```
heroku login
heroku git:remote -a <app>
-git remote rename heroku heroku-node-express
heroku buildpacks:add -a <app> https://github.com/lstoll/heroku-buildpack-monorepo
heroku buildpacks:add -a <app> heroku/nodejs
heroku config:set APP_BASE=heroku-node-express -a <app>
heroku addons:create heroku-postgresql:hobby-dev -a <app>
```
