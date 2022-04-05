# webstack

A project to prototype web stacks: runtimes, frameworks, databases, and hosting platforms.

For each stack,
`SELECT 1` to a SQL database
using a database driver (no ORM),
and return a JSON payload of `{ "status": "ok" }`,
deployed to a modern hosting platform.

Unless otherwise stated, Postgres is the database.

Runtimes / frameworks:

* Crystal / [Kemal](https://kemalcr.com/)
* Deno / stdlib, [Oak](https://oakserver.github.io/oak/)
* Go / stdlib, [Gin](https://github.com/gin-gonic/gin)
* Node / [Express](https://expressjs.com/)
* Ruby / [Sinatra](http://sinatrarb.com/), X

The dashboard at <https://webstack.checklyhq.com/>
shows real-time latencies from
Northern California, Montreal, Ireland, Stockholm, and Sydney.
