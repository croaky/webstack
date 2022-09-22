package main

import (
	"net"
	"net/url"
	"testing"
)

func TestPortReplacement(t *testing.T) {
	dbUrl := "postgres://user:pass@subdomain.host.com:5432/db_name"
	u, err := url.Parse(dbUrl)
	if err != nil {
		t.Fatalf(`could not parse %s`, dbUrl)
	}

	host, _, err := net.SplitHostPort(u.Host)
	if err != nil {
		t.Fatalf(`could not split host port %s`, u.Host)
	}
	u.Host = host + ":5433"
	got := u.String()

	want := "postgres://user:pass@subdomain.host.com:5433/db_name"

	if got != want {
		t.Fatalf(`got %s, want %s`, got, want)
	}
}
