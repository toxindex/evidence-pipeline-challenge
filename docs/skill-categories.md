# Skill categories

Choose one primary category. Build an end-to-end product, but spend your depth
budget where you want to be evaluated—and tell us which aspect you want to be
grilled on in the interview.

## 1. Scientific informatics

For people who want to demonstrate informatics and scientific judgment:

- a versioned BioBrick or equivalent dataset with separate retrieval,
  transformation and health checks;
- chemical standardization, identifier resolution, units, assays, sequences,
  ontologies or evidence integration;
- source/version tracking, idempotency, mapping-quality metrics and deliberate
  schema-drift failure.

We may ask you to explain a normalization decision, resolve an ambiguous entity,
replay an input, or trace a derived row to its source.

## 2. Data platforms & orchestration

For people who want to demonstrate distributed data and systems engineering:

- Spark/PySpark batch or structured streaming with schemas, checkpoints,
  partitioning, duplicate/late-event behavior and a bounded correctness test;
- a Temporal workflow with deterministic orchestration, idempotent activities,
  explicit retries and recovery after interruption;
- Kubernetes and Helm with probes, resource bounds, least privilege, secret
  references and a tested upgrade/rollback story;
- observability, backpressure, caching, security boundaries and honest partial
  failure.

We may ask you to diagnose skew, explain event-time behavior, interrupt a
workflow, or walk through deployment and rollback.

## 3. Intelligent applications

For people who want to turn evidence into a usable product with applied ML and
intelligent systems:

- a coherent API and human-facing interface designed from a real user's
  question, with an explicit contract and honest failure behavior;
- a defensible prediction, ranking, anomaly-detection or representation task
  with a baseline, leakage-safe split, relevant metrics and model card;
- deep learning where it materially beats or enables something the baseline
  cannot;
- an evidence-grounded agent with narrow tools, exact citations, bounded cost,
  validated output and adversarial/untrusted-input handling;
- uncertainty, applicability domain, monitoring and a demonstrated case where
  verification rejects an unsupported result.

We may ask you to explain an API or interaction, defend an evaluation, inspect a
failure, distinguish model signal from leakage, or show how an agent is
prevented from inventing evidence.
