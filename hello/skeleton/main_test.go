package main

import (
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGreeting(t *testing.T) {
	srv := httptest.NewServer(handler())
	defer srv.Close()

	res, err := http.Get(srv.URL + "/")
	if err != nil {
		t.Fatal(err)
	}
	defer res.Body.Close()
	body, _ := io.ReadAll(res.Body)
	if res.StatusCode != http.StatusOK || string(body) != "{% .Params.greeting %} from {% .Service.Name %}\n" {
		t.Fatalf("GET / = %d %q", res.StatusCode, body)
	}
}

func TestHealth(t *testing.T) {
	srv := httptest.NewServer(handler())
	defer srv.Close()

	res, err := http.Get(srv.URL + "/healthz")
	if err != nil {
		t.Fatal(err)
	}
	res.Body.Close()
	if res.StatusCode != http.StatusOK {
		t.Fatalf("GET /healthz = %d", res.StatusCode)
	}
}

func TestListenAddr(t *testing.T) {
	env := map[string]string{}
	getenv := func(k string) string { return env[k] }
	if got := listenAddr(getenv); got != ":{% .Params.port %}" {
		t.Fatalf("without PORT: %q", got)
	}
	env["PORT"] = "9000"
	if got := listenAddr(getenv); got != ":9000" {
		t.Fatalf("with PORT=9000: %q", got)
	}
}
