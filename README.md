# ToxIndex Evidence Pipeline Challenge

![Real-world evidence flowing through the capabilities ToxIndex is hiring for](assets/hiring-flow.svg)

Build a small, useful evidence product from public scientific, regulatory,
health, agricultural, or environmental data. Pick a real question, ingest at
least one real source, and ship a reproducible vertical slice with a documented
API and human-facing view.

## Choose your strength

Everyone builds the same core: trustworthy data, provenance from result back to
source, meaningful tests, and a product another person can use. Then choose
**one category** in which to show us real depth:

1. **Scientific informatics** — advanced ETL or a BioBrick, chemical and
   biological identifiers, scientific data quality, modeling and provenance.
2. **Data platforms & orchestration** — Spark/PySpark, streaming, Temporal,
   Kubernetes/Helm, reliability, observability and operability at scale.
3. **Intelligent applications** — useful APIs and interfaces, machine learning,
   deep learning, grounded agents, evaluation and verification.

See [skill categories](docs/skill-categories.md) for expectations and examples.
Technologies are evidence, not requirements: use only what improves the system.

## Build

- Choose a problem with real value to pharma, energy, cosmetics, agriculture,
  public health, regulators, or scientists reducing animal testing.
- Keep source identity, retrieval metadata and enough provenance to audit every
  result. Detect malformed or changed input; make reruns converge.
- Provide one-command setup, a small credential-free fixture, tests, an API
  contract, and a brief decision/limitations note.
- Aim for **6–8 focused hours**. Stop at the time box and tell us what comes next.

Start with the equivalent [Go](examples/go-openfda) or
[Python](examples/python-openfda) openFDA examples, browse
[self-hosted project gallery](docs/gallery.md) and
[public data sources](docs/data-sources.md), or start from scratch. We may also
provide [ToxIndex platform access](docs/toxindex.md) for optional enrichment.

## Agents are expected

We know everybody can code with agents today. Use them—we expect stronger,
more ambitious and better-tested submissions as a result. Disclose which tools
you used, where they helped, and what you personally verified. You remain
responsible for every design choice, claim and line you submit.

## Submit and discuss

Send a repository with the running system and a short README covering setup,
architecture, your chosen category, agent use, tradeoffs and limitations. Do
not commit credentials or improperly redistribute data.

We score against the public [rubric](RUBRIC.md). A submission that clears review
leads to an interview where you will run the system, explain how it works, trace
a result to its source, and reason through a failure. **You choose the aspect of
your selected skill category on which you want the deepest technical grilling.**
We will still discuss the whole system, but you decide where to set the bar.
