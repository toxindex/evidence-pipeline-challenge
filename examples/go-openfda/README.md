# Minimal Go openFDA ingestion example

This is the Go equivalent of the Python starter. It fetches recent drug recalls,
validates a small contract, writes NDJSON, and records retrieval provenance
using only the standard library.

```bash
go run . -limit 10 -output out/recalls.ndjson
go test ./...
```

It is not a complete submission: add the product, API, persistence, stronger
contracts, failure behavior, and the depth you want us to evaluate.

