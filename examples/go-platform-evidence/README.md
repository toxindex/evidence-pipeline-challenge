# Go ToxIndex evidence example

Discover Kiln/BioBricks datasets and live prediction tools, search source
evidence, run a predictor, and write a provenance-bearing evidence bundle.

```bash
export TOXINDEX_API_KEY=tidx_...
go run . -query "acetaminophen liver" -tool admet-ai \
  -smiles 'CC(=O)NC1=CC=C(C=C1)O' -output out/evidence.json
go test ./...
```

This is a client seam, not a submission. Build an API and interface around it;
persist exact source/tool/version metadata; then add orchestration, Spark, Helm,
or model/agent evaluation according to the category you choose.
