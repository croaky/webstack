# webstack

A project to prototype web stacks: frameworks, runtimes, databases, and hosting platforms.

For each stack,
`SELECT 1` to a SQL database
using a database driver (no ORM),
and return a JSON payload of `{ "status": "ok" }`,
deployed to a modern hosting platform.

Unless otherwise stated, Postgres is the database.

Frameworks (runtimes):

* [Sinatra (Ruby)](http://sinatrarb.com/)
* [Gin (Go)](https://github.com/gin-gonic/gin)
* [Express (Node)](https://expressjs.com/)
* [Oak (Deno)](https://oakserver.github.io/oak/)
* [Kemal (Crystal)](https://kemalcr.com/)

The dashboard at <https://webstack.checklyhq.com/>
shows real-time latencies from
Northern California, Montreal, Ireland, Stockholm, and Sydney.
