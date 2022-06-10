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
    @pool = ConnectionPool.new(size: 5, timeout: 5) {
      PG.connect(ENV.fetch("DATABASE_URL"))
    }
  end

  def do_exec(sql)
    @pool.with do |conn|
      conn.exec(sql)
    end
  end
end

db = DB.new

get "/" do
  db.exec "SELECT 1"
  content_type :json
  {status: "ok"}.to_json
end
