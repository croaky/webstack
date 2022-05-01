require_relative "config"
require_relative "../api"

describe "/health" do
  # https://inadarei.github.io/rfc-healthcheck/
  it "follows latest Internet-Draft" do
    resp = get "/health"

    expect(resp.status).to eq 200
    expect(resp.headers).to eq(
      "Accept" => "application/json",
      "Access-Control-Allow-Methods" => "GET",
      "Content-Length" => "15",
      "Content-Type" => "application/json",
      "X-Content-Type-Options" => "nosniff"
    )
    expect(resp.body).to eq %({"status":"ok"})
  end
end
