# Self-hosted gallery

These public projects show different shapes a self-hosted submission can take.
They are inspiration, not reference solutions or a hidden checklist. Several are
experiments rather than validated scientific products; read their limitations.

| | Project | Shape | Primary skill category |
|---|---|---|---|
| 🧬 | **[GeneScore](https://github.com/toxindex/genescore)** | Go + HTMX genomics application: FASTQ through established bioinformatics tools to a provenance-rich report | **Data & domain** |
| 🕸️ | **[Rhizome](https://github.com/toxindex/rhizome)** | Docker Compose toxicology knowledge graph with schemas, loaders, sample data and graph queries | **Data & domain** |
| 🧪 | **[PKPD](https://github.com/toxindex/PKPD)** | Containerized Python PK/PD and toxicity-modeling platform with CLI, notebooks and API-oriented deployment | **Models & agents** |
| 🖥️ | **[Box](https://github.com/toxindex/box)** | Single-binary Go + HTMX application, embedded assets, Docker image and air-gapped deployment posture | **Systems & product** |
| 🐹 | **[Go openFDA starter](../examples/go-openfda)** | Dependency-free live API ingestion, validation, NDJSON and retrieval provenance | **Starting point** |
| 🐍 | **[Python openFDA starter](../examples/python-openfda)** | The same ingestion contract in Python, making language choice irrelevant to the exercise | **Starting point** |

## What to notice

### GeneScore — a pipeline becomes a product

```text
FASTQ → alignment → variants → ClinVar / PGS Catalog → inspectable report
```

The interesting part is not merely invoking tools. It is job handling,
provenance, privacy, benchmark evidence, and communicating that the output is
educational—not clinical or diagnostic. The quick demo runs as a Go web app;
the full pipeline adds standard bioinformatics tools and larger references.

### Rhizome — evidence becomes connected

```text
source loaders → normalized entities and claims → graph → traversals
```

Rhizome shows a self-hosted graph stack with Docker Compose, an explicit schema,
sample loader, and questions that cross chemicals, targets, pathways, claims and
predictions. Look for identity strategy, provenance edges, conflict modeling,
and what happens when two sources disagree.

### PKPD — a model becomes an operable service

```text
chemical structures + observations → model / simulation → API, CLI and reports
```

PKPD illustrates a Python scientific stack packaged for local Docker use. A
challenge submission in this shape should go beyond producing predictions:
establish a baseline, explain applicability and uncertainty, validate inputs,
and make model/version provenance visible.

### Box — deployment constraints shape design

```text
embedded UI + server + data → one Go binary → local or air-gapped container
```

Box is a useful counterexample to reflexive distributed architecture. Its
self-hosted, air-gapped goal leads to embedded assets and a simple container.
Notice how a deployment constraint can justify fewer moving parts while still
leaving room for good APIs, tests, observability and security.

## Borrow the shape, not the answer

A strong submission may combine shapes—for example, a Rhizome-like evidence
graph served with Box-like operational simplicity, or a GeneScore-like pipeline
with a carefully evaluated model. It may also look nothing like these.

Whatever you choose, a reviewer should be able to clone it, start it locally,
follow one result back to source data, and discuss a failure with you.

