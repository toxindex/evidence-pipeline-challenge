# Review rubric

The core is scored out of 80. The candidate's chosen skill category contributes
up to 20. A submission is not penalized for omitting technologies outside it.

| Area | Points | Strong evidence |
|---|---:|---|
| Problem and user value | 10 | A concrete user and decision; the demo answers a real question |
| Data correctness | 20 | Explicit schema, normalization, validation, duplicate handling, meaningful tests |
| Provenance and reproducibility | 15 | Results trace to source records; inputs are versioned/checksummed; reruns converge |
| API and application | 15 | Usable vertical slice, coherent contract, honest errors, observable health |
| Reliability and security | 10 | Failure/retry behavior, bounded resources, secret handling, unsafe input considered |
| Scientific judgment | 10 | Identifiers/units/context handled; limitations and uncertainty clearly stated |
| Chosen skill category | 20 | Technically substantive, justified, tested, and integrated with the product |

## Review probes

We expect to ask the system to:

- ingest the same input twice without doubling results;
- encounter a malformed record, timeout, or upstream schema change;
- trace one displayed claim to the exact source and retrieval;
- explain an API response and a design tradeoff;
- demonstrate the chosen category under a realistic failure or edge case.

Candidates disclose their agent use and remain responsible for the submission.
Passing submissions proceed to an interview: the candidate runs and explains
the system, and chooses one aspect of their selected category for the deepest
technical discussion.

## What does not earn points

- architecture added only to satisfy a keyword;
- generated code the candidate cannot explain;
- a dashboard whose numbers cannot be traced to source records;
- model accuracy without a baseline, split rationale, or leakage analysis;
- an agent that presents unsupported prose as evidence;
- a Helm chart that has never been installed.
