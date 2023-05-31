# frozen_string_literal: true

require "connection_pool"
require "pg"
require "sinatra"

class DB
  def initialize
    @pool = ConnectionPool.new(size: 5, timeout: 5) {
      PG.connect(ENV.fetch("DATABASE_URL", "postgres:///webstack_dev"))
    }
  end

  def exec(sql)
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
