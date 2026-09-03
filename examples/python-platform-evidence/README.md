# Python ToxIndex evidence example

Discover Kiln/BioBricks datasets and live prediction tools, search source
evidence, run a predictor, and write a provenance-bearing evidence bundle using
only the Python standard library.

```bash
export TOXINDEX_API_KEY=tidx_...
python -m src.ingest --query "acetaminophen liver" --tool admet-ai \
  --smiles 'CC(=O)NC1=CC=C(C=C1)O' --output out/evidence.json
python -m unittest discover -s tests
```

This is a client seam, not a submission. Build the self-hosted product and the
depth you want us to evaluate. Never commit the platform key.
