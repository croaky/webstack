require_relative "lib/x"

X::Env.load(root_dir: __dir__)

X::API.serve do |db|
  get "/" do |req, resp|
    db.exec "SELECT 1"
    [
      200,
      resp.headers,
      {status: "ok"}.to_json
    ]
  end
end
