require "connection_pool"
require "pg"
require "sinatra"

class DB
  def initialize
    connect
  end

  def exec(sql)
    do_exec(sql)
  rescue PG::ConnectionBad
    connect
    do_exec(sql)
  end

  private

  def connect
    url = ENV.fetch("DATABASE_URL")
    primary = ENV["PRIMARY_REGION"].to_s
    current = ENV["FLY_REGION"].to_s

    if primary != "" && current != "" && primary != current
      u = URI.parse(url)
      u.port = 5433
      url = u.to_s
    end

    @pool = ConnectionPool.new(size: 5, timeout: 5) {
      PG.connect(url)
    }
  end

  def do_exec(sql)
    @pool.with do |conn|
      conn.exec(sql)
    end
  end
end

db = DB.new

configure do
  set :protection, except: [:json_csrf]
end

get "/" do
  db.exec "SELECT 1"
  content_type :json
  {status: "ok"}.to_json
end
