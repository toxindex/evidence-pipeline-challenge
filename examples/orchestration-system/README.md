# Example A — Evidence operations system

Primary category: **Data platforms & orchestration**  
Supporting skills: API design, identity, observability, scientific provenance

Build a self-hosted system that refreshes evidence for a watchlist of chemicals,
queries licensed datasets in ToxIndex Kiln, runs relevant ToxIndex predictors,
and publishes an auditable API/dashboard. The workload must survive retries,
partial upstream failure, deployments, and restarts without duplicating results.

This is not “put one script in five frameworks.” Give each component a reason
to exist and say what you would remove at smaller scale.

## Required system

```text
schedule / API
      │
      ▼
Temporal workflow ──▶ Kiln brick/search/SPARQL APIs
      │
      ├─────────────▶ ToxIndex prediction tools
      │
      ▼
Spark transform ──▶ versioned evidence store ──▶ API + dashboard
      │
      └── metrics, lineage, dead letters and replay
```

- **Temporal** owns the durable per-chemical workflow: resolve inputs, fetch
  evidence, submit predictions, wait/retry, publish atomically. Activities are
  idempotent; retryable and terminal errors differ; workflow code can replay.
- **Spark** performs a justified cross-source or incremental computation—for
  example normalizing a Tox21/ToxCast evidence matrix, deduplicating by source
  identity, or computing watchlist deltas. Show partitions, event time,
  checkpointing, skew/duplicate behavior, and a bounded correctness fixture.
- **Airflow** schedules/backfills dataset-level runs and reports data quality.
  It must not compete with Temporal for ownership of the same retry state.
  Explain the boundary, or omit Airflow and defend that choice.
- **Kubernetes + Helm** run the system locally in kind/k3d. Include probes,
  resource requests/limits, jobs/migrations, secret references, network policy,
  service account, persistent-state behavior, and a tested upgrade/rollback.
- **Application API** exposes run status, evidence, prediction provenance and
  failure state. The UI shows freshness and partial results rather than hiding
  them.

## ToxIndex integration and identity

Use `GET /v1/kiln/brick`, Kiln search or SPARQL, prediction-tool discovery, and
`POST /v1/tools/toxjobs_predict`. Discover schemas at runtime and preserve the
brick source/license/build time and prediction tool metadata used for each run.

Interactive local login should use the platform's OAuth device grant. Workloads
use a Kubernetes Secret containing a supplied `tidx_` key. Explain why browser
OIDC, workload identity and an API key are different credentials; never bake a
token into an image, Helm value, log or Spark plan.

## Failure review

We will interrupt a workflow, repeat an event, make one prediction unavailable,
rotate a credential, and upgrade the Helm release. We should see recovery or an
honest terminal state, never silent loss or duplicate publication.

## What to submit

- one-command local boot and teardown;
- Temporal workflow plus replay/idempotency tests;
- Spark job and bounded fixture with expected output;
- Airflow DAG or a short architectural decision explaining its omission;
- chart, values schema, rendered-manifest test and rollback notes;
- OpenAPI-described service and small status/evidence UI;
- trace/metrics screenshot and a runbook for one injected failure.

## Deliberately flawed starter

This directory is intentionally incomplete. Its sample DAG, Spark job and Helm
chart contain reviewable faults: overlapping orchestration ownership, an
unbounded collect, mutable image tags, a literal credential, weak probes and no
idempotency key. Do not patch mechanically—replace the design where necessary
and record what you found.

