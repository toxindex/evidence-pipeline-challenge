package main

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"strings"
	"time"
)

type platform struct {
	base, key string
	http      *http.Client
}

func newPlatform(key, base string) (*platform, error) {
	if key == "" {
		return nil, errors.New("TOXINDEX_API_KEY is required")
	}
	if base == "" {
		base = "https://platform.toxindex.com"
	}
	return &platform{strings.TrimRight(base, "/"), key, &http.Client{Timeout: 120 * time.Second}}, nil
}

func (p *platform) call(ctx context.Context, method, route string, body any) (json.RawMessage, error) {
	var data []byte
	var err error
	if body != nil {
		data, err = json.Marshal(body)
		if err != nil {
			return nil, err
		}
	}
	req, err := http.NewRequestWithContext(ctx, method, p.base+route, bytes.NewReader(data))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Authorization", "Bearer "+p.key)
	if body != nil {
		req.Header.Set("Content-Type", "application/json")
	}
	res, err := p.http.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	raw, err := io.ReadAll(io.LimitReader(res.Body, 16<<20))
	if err != nil {
		return nil, err
	}
	if res.StatusCode >= 400 {
		return nil, fmt.Errorf("platform %s: %s", res.Status, strings.TrimSpace(string(raw)))
	}
	if !json.Valid(raw) {
		return nil, errors.New("platform returned invalid JSON")
	}
	return raw, nil
}

func buildBundle(ctx context.Context, p *platform, query, tool, smiles string) (map[string]any, error) {
	get := func(route string) (json.RawMessage, error) { return p.call(ctx, http.MethodGet, route, nil) }
	bricks, err := get("/v1/kiln/brick")
	if err != nil {
		return nil, err
	}
	search, err := get("/v1/kiln/search?" + url.Values{"q": {query}, "limit": {"10"}}.Encode())
	if err != nil {
		return nil, err
	}
	tools, err := get("/v1/toxjobs/tool")
	if err != nil {
		return nil, err
	}
	prediction, err := p.call(ctx, http.MethodPost, "/v1/tools/toxjobs_predict", map[string]string{"tool": tool, "smiles": smiles})
	if err != nil {
		return nil, err
	}
	return map[string]any{
		"schema": "toxindex-evidence-bundle/v1", "created_at": time.Now().UTC().Format(time.RFC3339Nano),
		"query": query, "input": map[string]string{"smiles": smiles},
		"sources":    map[string]json.RawMessage{"bricks": bricks, "search": search},
		"prediction": map[string]any{"requested_tool": tool, "catalog_at_run": tools, "result": prediction},
	}, nil
}

func main() {
	query := flag.String("query", "", "evidence search query")
	tool := flag.String("tool", "", "prediction tool name")
	smiles := flag.String("smiles", "", "molecule as SMILES")
	output := flag.String("output", "out/evidence.json", "output path")
	flag.Parse()
	if *query == "" || *tool == "" || *smiles == "" {
		flag.Usage()
		os.Exit(2)
	}
	p, err := newPlatform(os.Getenv("TOXINDEX_API_KEY"), os.Getenv("TOXINDEX_BASE_URL"))
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	bundle, err := buildBundle(context.Background(), p, *query, *tool, *smiles)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	if err := os.MkdirAll(filepath.Dir(*output), 0o755); err != nil {
		panic(err)
	}
	raw, _ := json.MarshalIndent(bundle, "", "  ")
	if err := os.WriteFile(*output, append(raw, '\n'), 0o644); err != nil {
		panic(err)
	}
	fmt.Println("wrote evidence bundle to", *output)
}
