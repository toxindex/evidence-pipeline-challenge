# Optional ToxIndex resources

ToxIndex access is issued for the challenge. The submission must still have a
credential-free fixture so reviewers can reproduce its core behavior offline.

The public Go SDK is [`github.com/toxindex/toxindex-go`](https://github.com/toxindex/toxindex-go).
It provides:

- a machine-plane client for `platform.toxindex.com/v1`;
- an MCP endpoint at `https://platform.toxindex.com/mcp` for coding agents;
- gateway clients for chemical-property predictions, literature, regulatory
  documents, identifier resolution, and other open-access resources.

The live API has a public OpenAPI document at
[`/v1/openapi.json`](https://platform.toxindex.com/v1/openapi.json). Useful
machine endpoints include:

- `GET /v1/kiln/brick` — discover curated Kiln/BioBricks datasets;
- `GET /v1/kiln/search` and `GET|POST /v1/kiln/sparql` — retrieve source
  evidence with brick/document provenance;
- `GET /v1/toxjobs/tool` — discover prediction tools currently served by ready
  rigs;
- `GET /v1/tools` and `POST /v1/tools/{name}` — discover JSON Schemas and call
  tools such as `toxjobs_predict`;
- `POST /mcp` — mount the same tool surface in an agent.

Platform keys begin with `tidx_` and should be read from `TOXINDEX_API_KEY`.
Never commit a key. The API is still evolving, so discover tool schemas at
runtime and keep captured fixtures for reproducible tests. Treat predictions as
derived evidence: retain tool/version metadata and link results back to the
underlying public source rows where available.

Using ToxIndex does not automatically satisfy the agent, ML, or domain track.
We review what you built around the resource and how you validated it.
