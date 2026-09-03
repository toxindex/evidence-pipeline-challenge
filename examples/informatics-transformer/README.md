# Example B — Tox21 molecular representation lab

Primary category: **Scientific informatics**  
Supporting skills: data processing, cheminformatics, ML/deep learning, product

Build a reproducible molecular representation and exploration service from the
real Tox21 multi-label assay dataset in BioBricks/Kiln. Train a small transformer
or other learned molecular encoder, cluster its embeddings, and test whether the
representation captures chemically and toxicologically meaningful structure.
Enrich selected compounds with live ToxIndex predictions and source evidence.

The goal is not “train a transformer.” It is to show that you understand the
chemical records, the assay labels, the evaluation traps, and what the learned
space is useful for.

## Required analysis

1. Discover the Tox21 brick and table schema from the ToxIndex API; export a
   pinned training snapshot or reproducible query with source/license/build
   metadata. EPA ToxCast is a strong scale-up or external-validation source.
2. Standardize structures deliberately: invalid SMILES, salts, mixtures,
   stereochemistry, tautomer policy, duplicates and conflicting labels must be
   measured—not silently discarded.
3. Establish non-neural baselines (for example Morgan fingerprints with a
   linear/tree model, and fingerprint clustering).
4. Train or fine-tune a SMILES/graph transformer for multi-label endpoints or
   representation learning. Use scaffold-aware splits and compare against the
   baseline; random-only splits are insufficient.
5. Cluster embeddings with a method appropriate to the geometry. Report
   stability, chemical diversity and endpoint enrichment; do not treat a UMAP
   picture as validation.
6. For representative clusters and outliers, call ToxIndex predictors and query
   Kiln evidence. Keep measured assay labels, computed predictions and model
   embeddings as distinct provenance classes.
7. Serve a small application: search a molecule, inspect neighbors/clusters,
   compare measured and predicted endpoints, and trace every datum to source.

## What we will inspect

- leakage through duplicate structures, scaffolds or preprocessing fitted on
  the full dataset;
- missing-label masking and class imbalance in multi-task evaluation;
- molecule identity and unit/assay semantics;
- baseline parity, ablations, calibration and uncertainty;
- reproducible environment, seeds, model/data cards and artifact hashes;
- whether a cluster claim survives quantitative inspection.

## What to submit

- versioned extraction and validation report;
- baseline and transformer training/evaluation code;
- embedding/clustering analysis with stability metrics;
- model card and scientific limitations;
- API plus interactive local explorer;
- CPU-sized fixture/test path; full training may use a GPU.

## Deliberately flawed starter

The starter uses a random row split, fits preprocessing before splitting,
conflates missing labels with negatives, and chooses cluster count from the same
visualization it reports. It is an interview artifact: identify, measure and fix
the problems rather than assuming the scaffold is authoritative.

