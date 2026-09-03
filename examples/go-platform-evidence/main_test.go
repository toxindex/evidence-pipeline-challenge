package main

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestBundleKeepsSourcesAndPredictionCatalog(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc("/v1/kiln/brick", func(w http.ResponseWriter, _ *http.Request) { w.Write([]byte(`{"items":[{"slug":"tox21"}]}`)) })
	mux.HandleFunc("/v1/kiln/search", func(w http.ResponseWriter, _ *http.Request) { w.Write([]byte(`{"hits":[{"brick":"tox21"}]}`)) })
	mux.HandleFunc("/v1/toxjobs/tool", func(w http.ResponseWriter, _ *http.Request) { w.Write([]byte(`{"items":[{"name":"admet-ai"}]}`)) })
	mux.HandleFunc("/v1/tools/toxjobs_predict", func(w http.ResponseWriter, _ *http.Request) { w.Write([]byte(`{"score":0.4}`)) })
	srv := httptest.NewServer(mux)
	defer srv.Close()
	p, _ := newPlatform("tidx_test", srv.URL)
	bundle, err := buildBundle(context.Background(), p, "liver", "admet-ai", "CCO")
	if err != nil {
		t.Fatal(err)
	}
	sources := bundle["sources"].(map[string]json.RawMessage)
	if !bytesContain(sources["search"], `"brick":"tox21"`) {
		t.Fatalf("source provenance missing: %s", sources["search"])
	}
	pred := bundle["prediction"].(map[string]any)
	if pred["requested_tool"] != "admet-ai" {
		t.Fatalf("tool provenance missing: %#v", pred)
	}
}

func TestRequiresKey(t *testing.T) {
	if _, err := newPlatform("", ""); err == nil {
		t.Fatal("expected missing key error")
	}
}

func bytesContain(got json.RawMessage, want string) bool { return strings.Contains(string(got), want) }
