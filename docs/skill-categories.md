# Skill categories

Choose one primary category. Build an end-to-end product, but spend your depth
budget where you want to be evaluated—and tell us which aspect you want to be
grilled on in the interview.

## 1. Data & domain

For people who want to demonstrate data engineering and scientific judgment:

- a versioned BioBrick or equivalent dataset with separate retrieval,
  transformation and health checks;
- Spark/PySpark batch or structured streaming with schemas, checkpoints,
  duplicate/late-event behavior and a bounded correctness test;
- chemical standardization, identifier resolution, units, assays, sequences,
  ontologies or evidence integration;
- source/version tracking, idempotency, mapping-quality metrics and deliberate
  schema-drift failure.

We may ask you to explain a normalization decision, replay an input, diagnose
skew or duplicates, or trace a derived row to its source.

## 2. Systems & product

For people who want to demonstrate application engineering and operations:

- a coherent API and useful interface designed from a real user's question;
- a Temporal workflow with deterministic orchestration, idempotent activities,
  explicit retries and recovery after interruption;
- Kubernetes and Helm with probes, resource bounds, least privilege, secret
  references and a tested upgrade/rollback story;
- observability, backpressure, caching, security boundaries and honest partial
  failure.

We may ask you to explain a contract, interrupt a workflow, diagnose a failing
request, or walk through deployment and rollback.

## 3. Models & agents

For people who want to demonstrate applied ML and intelligent systems:

- a defensible prediction, ranking, anomaly-detection or representation task
  with a baseline, leakage-safe split, relevant metrics and model card;
- deep learning where it materially beats or enables something the baseline
  cannot;
- an evidence-grounded agent with narrow tools, exact citations, bounded cost,
  validated output and adversarial/untrusted-input handling;
- uncertainty, applicability domain, monitoring and a demonstrated case where
  verification rejects an unsupported result.

We may ask you to defend an evaluation, inspect a failure, distinguish model
signal from leakage, or show how an agent is prevented from inventing evidence.

