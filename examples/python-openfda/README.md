# Minimal openFDA ingestion example

This example fetches recent drug recall records, validates a deliberately small
contract, writes newline-delimited JSON, and records retrieval provenance. It
uses only the Python standard library.

```bash
python -m src.ingest --limit 10 --output out/recalls.ndjson
python -m unittest discover -s tests
```

It is starter code, not a complete submission. It lacks persistence, an API,
a UI, incremental checkpoints, robust schema-drift handling, and most fields a
real analysis would need. Improve it or replace it.

openFDA permits limited unauthenticated use. Set `OPENFDA_API_KEY` if you have a
key; never commit it.

