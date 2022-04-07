# webstack

A project to prototype web stacks.

For each stack,
`SELECT 1` to a Postgres database
using a database driver (no ORM)
and return a JSON payload of `{ "status": "ok" }`.

Runtimes / frameworks:

* Crystal / [Kemal](https://kemalcr.com/)
* Deno / [stdlib](https://deno.land/std)
* Go / [stdlib](https://pkg.go.dev/std)
* Node / [Express](https://expressjs.com/)
* Ruby / [Sinatra](http://sinatrarb.com/)

Hosting platforms:

* [Deno Deploy](https://deno.com/deploy)
* [Fly.io](https://fly.io)
* [Heroku](https://heroku.com)
* [Railway](https://railway.app)
* [Render](https://render.com)
* [Vercel](https://vercel.com)

The dashboard at <https://webstack.checklyhq.com/>
shows real-time request data from
N. California, Montreal, Ireland, Stockholm, and Sydney.
