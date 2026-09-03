# Starter examples

The Go and Python examples perform the same useful operation against
`platform.toxindex.com`: discover available Kiln/BioBricks datasets, search
their evidence, discover live prediction tools, and run a selected predictor.
They emit one provenance-bearing evidence bundle rather than pretending a model
result and source passage are the same kind of evidence.

- [`go-platform-evidence`](go-platform-evidence)
- [`python-platform-evidence`](python-platform-evidence)

Both use only their language's standard library and include fixture-based unit
tests. A `tidx_` API key is required for live calls.

The three full challenge examples are deliberately flawed systems with detailed
candidate briefs:

- [`orchestration-system`](orchestration-system) — Temporal, Spark, Airflow,
  Kubernetes/Helm, auth and a managed ToxIndex evidence application.
- [`informatics-transformer`](informatics-transformer) — Tox21 molecular
  representation learning, clustering, cheminformatics and ToxIndex enrichment.
- [`agentic-reviewer`](agentic-reviewer) — a bounded, evaluated evidence agent
  over ToxIndex tools with citation and injection defenses.

Candidates may take one of these shapes or start over. The faults are explicit
enough to invite review, not a puzzle about undocumented broken code.
