require "kemal"

get "/" do |env|
  env.response.content_type = "application/json"
  {"status": "ok"}.to_json
end

Kemal.run
