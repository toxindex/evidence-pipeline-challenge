# ToxIndex Evidence Pipeline Challenge

Build a small, living evidence product from public scientific, regulatory,
health, agricultural, or environmental data.

This is deliberately open-ended. We want to see how you find a useful question,
turn unruly source data into a trustworthy product, and make technical choices
you can defend. A narrow system with excellent provenance and failure behavior
is stronger than a large diagram held together by happy-path demos.

## The brief

Choose a real-world question valuable to one of these groups:

- pharma or biotech researchers;
- energy, agriculture, chemical, or cosmetics safety teams;
- regulators or public-health organizations such as FDA, EPA, CDC, or NIEHS;
- scientists trying to replace or reduce animal testing.

Then build a vertical slice that:

1. ingests at least one real public data source;
2. preserves source identity, retrieval time, and enough provenance to audit a
   result;
3. produces a clean, typed data model and detects malformed or changed input;
4. exposes a documented API plus a useful human-facing view;
5. is reproducible from a clean checkout and tested around meaningful failure
   modes; and
6. explains what the evidence can—and cannot—support.

Choose **one or two depth tracks** from [the tracks](docs/tracks.md). Do not add
Kubernetes, Spark, Temporal, an LLM, or deep learning merely to name them. If a
single process is the right design, say so and make that process excellent.

## Time box

Aim for **6–8 focused hours**. Stop when the time box ends and describe the next
steps. We do not expect a production system or a polished research result.
Using coding assistants is welcome; disclose how you used them and what you
personally verified.

## Starting points

- [`examples/python-openfda`](examples/python-openfda) is a tiny runnable
  ingestion example. It is intentionally incomplete and not an expected
  architecture.
- [`docs/data-sources.md`](docs/data-sources.md) lists public live and static
  sources across drug safety, toxicology, environment, and public health.
- [`data/seed-substances.csv`](data/seed-substances.csv) provides a few chemical
  identifiers for quick experiments.
- ToxIndex API access may be provided. It is an optional enrichment surface,
  never the only source of truth; see [`docs/toxindex.md`](docs/toxindex.md).

You may build a BioBrick, a durable ETL workflow, a Spark stream and dashboard,
a chemical or biological model, an evidence-grounded agent, or something we
did not anticipate. The result should answer a real question—not just move
records between technologies.

## Deliverables

Submit a repository containing:

- the running application and pipeline;
- a README with one-command setup, architecture, chosen depth tracks, and a
  short demo path;
- the data contract/schema and provenance strategy;
- tests and a sample or fixture that run without private credentials;
- API documentation (OpenAPI is ideal, but a precise equivalent is fine);
- a short decision log: important tradeoffs, failure modes, security/privacy
  considerations, and what you would do next;
- if you trained a model, its baseline, held-out evaluation, and model card;
- if you deploy to Kubernetes, a Helm chart and local installation instructions.

Do not commit credentials or redistribute data whose license does not permit it.
Prefer download scripts plus checksums over unexplained copied datasets.

## How we review

The rubric is public in [`RUBRIC.md`](RUBRIC.md). We care most about judgment,
correctness, provenance, usability, and depth in the areas you chose. We will
clone the repo, run the documented happy path, inject a failure or duplicate,
inspect one result back to its source, and discuss your decisions with you.

Questions and reasonable assumptions are welcome. Record assumptions in the
repository so we can discuss them.

