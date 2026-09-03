package main

import (
	"context"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"
)

const endpoint = "https://api.fda.gov/drug/enforcement.json"

type recall struct {
	RecallNumber       string `json:"recall_number"`
	Status             string `json:"status"`
	Classification     string `json:"classification"`
	ProductDescription string `json:"product_description"`
	ReasonForRecall    string `json:"reason_for_recall"`
	ReportDate         string `json:"report_date,omitempty"`
	Source             string `json:"source"`
}

func normalize(raw map[string]any) (recall, error) {
	id, _ := raw["recall_number"].(string)
	id = strings.TrimSpace(id)
	if id == "" {
		return recall{}, errors.New("record has no recall_number")
	}
	str := func(key, fallback string) string {
		if value, ok := raw[key].(string); ok {
			return value
		}
		return fallback
	}
	return recall{RecallNumber: id, Status: str("status", "unknown"),
		Classification:     str("classification", "unknown"),
		ProductDescription: str("product_description", ""),
		ReasonForRecall:    str("reason_for_recall", ""),
		ReportDate:         str("report_date", ""), Source: endpoint}, nil
}

func fetch(ctx context.Context, client *http.Client, limit int, apiKey string) ([]recall, []byte, string, error) {
	params := url.Values{"limit": {strconv.Itoa(limit)}, "sort": {"report_date:desc"}}
	if apiKey != "" {
		params.Set("api_key", apiKey)
	}
	requestURL := endpoint + "?" + params.Encode()
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, requestURL, nil)
	if err != nil {
		return nil, nil, "", err
	}
	req.Header.Set("User-Agent", "toxindex-evidence-challenge/1.0")
	res, err := client.Do(req)
	if err != nil {
		return nil, nil, "", err
	}
	defer res.Body.Close()
	if res.StatusCode != http.StatusOK {
		return nil, nil, "", fmt.Errorf("openFDA returned %s", res.Status)
	}
	body, err := io.ReadAll(io.LimitReader(res.Body, 16<<20))
	if err != nil {
		return nil, nil, "", err
	}
	var payload struct {
		Results []map[string]any `json:"results"`
	}
	if err := json.Unmarshal(body, &payload); err != nil {
		return nil, nil, "", err
	}
	if payload.Results == nil {
		return nil, nil, "", errors.New("openFDA response has no results array")
	}
	records := make([]recall, 0, len(payload.Results))
	for _, item := range payload.Results {
		record, err := normalize(item)
		if err != nil {
			return nil, nil, "", err
		}
		records = append(records, record)
	}
	return records, body, requestURL, nil
}

func write(records []recall, raw []byte, requestURL, apiKey, output string) error {
	if err := os.MkdirAll(filepath.Dir(output), 0o755); err != nil {
		return err
	}
	f, err := os.Create(output)
	if err != nil {
		return err
	}
	enc := json.NewEncoder(f)
	for _, record := range records {
		if err := enc.Encode(record); err != nil {
			f.Close()
			return err
		}
	}
	if err := f.Close(); err != nil {
		return err
	}
	if apiKey != "" {
		requestURL = strings.ReplaceAll(requestURL, apiKey, "REDACTED")
	}
	sum := sha256.Sum256(raw)
	provenance := map[string]any{"retrieved_at": time.Now().UTC().Format(time.RFC3339Nano),
		"request_url": requestURL, "raw_sha256": hex.EncodeToString(sum[:]), "records": len(records)}
	b, err := json.MarshalIndent(provenance, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(output+".provenance.json", append(b, '\n'), 0o644)
}

func main() {
	limit := flag.Int("limit", 10, "records to fetch (1-100)")
	output := flag.String("output", "out/recalls.ndjson", "output path")
	flag.Parse()
	if *limit < 1 || *limit > 100 {
		fmt.Fprintln(os.Stderr, "limit must be between 1 and 100")
		os.Exit(2)
	}
	apiKey := os.Getenv("OPENFDA_API_KEY")
	records, raw, requestURL, err := fetch(context.Background(), &http.Client{Timeout: 30 * time.Second}, *limit, apiKey)
	if err == nil {
		err = write(records, raw, requestURL, apiKey, *output)
	}
	if err != nil {
		fmt.Fprintln(os.Stderr, "ingest:", err)
		os.Exit(1)
	}
	fmt.Printf("wrote %d records to %s\n", len(records), *output)
}
