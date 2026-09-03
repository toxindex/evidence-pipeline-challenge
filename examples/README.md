# Starter examples

The Go and Python examples perform the same useful operation against
`platform.toxindex.com`: discover available Kiln/BioBricks datasets, search
their evidence, discover live prediction tools, and run a selected predictor.
They emit one provenance-bearing evidence bundle rather than pretending a model
result and source passage are the same kind of evidence.

- [`go-platform-evidence`](go-platform-evidence)
- [`python-platform-evidence`](python-platform-evidence)

Both use only their language's standard library and include fixture-based unit
tests. A `tidx_` API key is required for live calls. They are deliberately
incomplete: candidates should turn the client into a self-hosted product and
demonstrate depth in their chosen skill category.
