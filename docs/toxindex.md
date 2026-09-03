# Optional ToxIndex resources

ToxIndex access may be issued to a candidate for enrichment and exploration.
The submission must still have a credential-free fixture or public-source path
so reviewers can reproduce its core behavior.

The public Go SDK is [`github.com/toxindex/toxindex-go`](https://github.com/toxindex/toxindex-go).
It provides:

- a machine-plane client for `platform.toxindex.com/v1`;
- an MCP endpoint at `https://platform.toxindex.com/mcp` for coding agents;
- gateway clients for chemical-property predictions, literature, regulatory
  documents, identifier resolution, and other open-access resources.

Platform keys begin with `tidx_` and should be read from
`TOXINDEX_API_KEY`. Never commit a key. Treat ToxIndex responses as derived
evidence: retain tool/version metadata and link results back to their underlying
public sources where available.

Using ToxIndex does not automatically satisfy the agent, ML, or domain track.
We review what you built around the resource and how you validated it.

