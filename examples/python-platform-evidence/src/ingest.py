"""Build a small evidence bundle from ToxIndex Kiln and prediction tools."""

from __future__ import annotations

import argparse
import datetime as dt
import json
import os
import pathlib
import urllib.parse
import urllib.request
from typing import Any


class Platform:
    def __init__(self, key: str, base: str = "https://platform.toxindex.com") -> None:
        if not key:
            raise ValueError("TOXINDEX_API_KEY is required")
        self.key, self.base = key, base.rstrip("/")

    def request(self, method: str, route: str, body: Any = None) -> Any:
        data = None if body is None else json.dumps(body).encode()
        request = urllib.request.Request(self.base + route, data=data, method=method)
        request.add_header("Authorization", "Bearer " + self.key)
        if data is not None:
            request.add_header("Content-Type", "application/json")
        with urllib.request.urlopen(request, timeout=120) as response:
            return json.load(response)

    def bricks(self) -> Any:
        return self.request("GET", "/v1/kiln/brick")

    def search(self, query: str) -> Any:
        q = urllib.parse.urlencode({"q": query, "limit": 10})
        return self.request("GET", "/v1/kiln/search?" + q)

    def prediction_tools(self) -> Any:
        return self.request("GET", "/v1/toxjobs/tool")

    def predict(self, tool: str, smiles: str) -> Any:
        name = urllib.parse.quote("toxjobs_predict", safe="")
        return self.request("POST", "/v1/tools/" + name,
                            {"tool": tool, "smiles": smiles})


def build_bundle(api: Platform, query: str, tool: str, smiles: str) -> dict[str, Any]:
    """Keep source evidence, catalog state, and derived prediction distinct."""
    return {
        "schema": "toxindex-evidence-bundle/v1",
        "created_at": dt.datetime.now(dt.timezone.utc).isoformat(),
        "query": query,
        "input": {"smiles": smiles},
        "sources": {"bricks": api.bricks(), "search": api.search(query)},
        "prediction": {
            "requested_tool": tool,
            "catalog_at_run": api.prediction_tools(),
            "result": api.predict(tool, smiles),
        },
    }


def main() -> None:
    parser = argparse.ArgumentParser()
    parser.add_argument("--query", required=True)
    parser.add_argument("--tool", required=True)
    parser.add_argument("--smiles", required=True)
    parser.add_argument("--output", type=pathlib.Path, default=pathlib.Path("out/evidence.json"))
    args = parser.parse_args()
    bundle = build_bundle(Platform(os.getenv("TOXINDEX_API_KEY", "")),
                          args.query, args.tool, args.smiles)
    args.output.parent.mkdir(parents=True, exist_ok=True)
    args.output.write_text(json.dumps(bundle, indent=2) + "\n")
    print(f"wrote evidence bundle to {args.output}")


if __name__ == "__main__":
    main()

