# Example C — Evidence-grounded chemical review agent

Primary category: **Intelligent applications**  
Supporting skills: agent design, API/product engineering, informatics, evaluation

Build a self-hosted review workbench where an agent investigates a chemical or
safety question using ToxIndex tools. It should query Kiln bricks/documents or
SPARQL, resolve chemical identity, run appropriate prediction tools, and produce
a structured evidence brief whose claims are traceable to exact tool results.

The challenge is not calling an LLM. It is designing a bounded system that
knows the difference between source evidence, a computed prediction and its own
interpretation.

## Required agent

- Mount `https://platform.toxindex.com/mcp` or dynamically load `/v1/tools` JSON
  Schemas. Do not hard-code a fictional tool catalog.
- Give the model narrow tools for source discovery, evidence retrieval,
  predictions and citation resolution. Separate read tools from any action.
- Resolve the chemical before searching. Preserve alternate identifiers and
  make ambiguity visible.
- Require structured claims: statement, evidence IDs, evidence class
  (`measured`, `source-text`, `prediction`, `inference`), confidence and
  limitations. Reject citations that were not present in the tool trace.
- Bound rounds, tokens, concurrency and spend. Support cancellation, timeouts,
  partial answers and resumable conversation history.
- Treat retrieved documents as untrusted data. Demonstrate prompt-injection and
  tool-output attacks that fail safely.
- Expose the trace in a usable interface so a reviewer can open the exact Kiln
  passage/row and prediction metadata behind a claim.

## Evaluation set

Create 15–30 questions with expected evidence requirements, including ambiguous
chemical names, unsupported conclusions, conflicting sources, missing data and
a prediction that must not be presented as experimental fact. Score citation
validity, evidence-class accuracy, answer completeness, abstention and cost—not
just prose similarity. Include at least one deterministic verifier.

## What we will do

We will insert an instruction into a retrieved document, remove a tool during a
run, ask an unanswerable question, replay a saved trace, and inspect whether the
answer's citations and evidence classes are mechanically valid.

## What to submit

- self-hosted API/UI and one-command local boot;
- versioned system prompt, dynamic tool adapter and trace store;
- structured answer schema plus citation verifier;
- evaluation dataset, runner, metrics and failure analysis;
- threat model, cost/latency budget and one red-team transcript;
- explicit disclosure of which agent wrote which parts of the submission.

## Deliberately flawed starter

The starter trusts retrieved text, accepts model-supplied URLs as citations,
has no tool-round budget, mixes predictions with observations and evaluates only
whether an answer is non-empty. Replace those choices and explain why.

