# Depth tracks

Choose one or two. State why the track belongs in your design and how you tested
it. These are prompts, not rigid feature lists.

## Reproducible dataset / BioBrick

Package a source as a versioned, reusable data product. Separate retrieval,
transformation, and health checks; pin or checksum raw inputs; record license
and provenance; publish a typed artifact such as Parquet or RDF. Demonstrate
schema-drift detection and idempotent rebuilding.

## Live stream and distributed data

Continuously ingest a changing source and provide freshness/lag metrics. Spark
Structured Streaming or PySpark is welcome when justified. Show event-time
semantics, checkpointing, duplicate/late-event handling, partition strategy,
and a correctness test on a local bounded fixture.

## Durable orchestration with Temporal

Model ingestion and enrichment as a durable workflow. Make activity boundaries
and retry policy explicit; keep activities idempotent; distinguish retryable
and terminal failures. Demonstrate that an interrupted workflow resumes without
duplicating published output. Include a workflow replay test if practical.

## Kubernetes and Helm

Package the working system for a local Kubernetes cluster. The chart should
lint and install with documented values, probes, resource requests/limits,
least-privilege security contexts, and secret references. Explain upgrade,
rollback, job retry, and persistent-state behavior. A local kind or k3d demo is
enough; no paid cloud account is expected.

## Machine learning / deep learning

Pose a defensible prediction, ranking, anomaly-detection, or representation
problem. Establish a simple baseline, prevent train/test leakage, justify the
split and metrics, version inputs/model, quantify uncertainty, and supply a
model card. Deep learning should outperform or enable something the baseline
cannot—not merely add a framework.

## Bioinformatics / cheminformatics

Demonstrate domain judgment: chemical standardization, identifier mapping,
units and assay context, sequence/structure handling, ontologies, or evidence
integration. Preserve original records alongside normalized entities and
measure unresolved/ambiguous mappings.

## Evidence-grounded agent

Build an agent that answers or acts through narrow tools over your evidence
store. Treat retrieved text and model output as untrusted, cite exact records,
validate structured outputs, bound cost/iterations, and demonstrate a case where
verification rejects or corrects an unsupported answer.

