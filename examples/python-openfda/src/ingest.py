"""Small, dependency-free openFDA recall ingestor used as challenge starter."""

from __future__ import annotations

import argparse
import datetime as dt
import hashlib
import json
import os
import pathlib
import urllib.parse
import urllib.request
from typing import Any

ENDPOINT = "https://api.fda.gov/drug/enforcement.json"


def normalize(record: dict[str, Any]) -> dict[str, Any]:
    """Validate the fields we rely on while retaining the source identifier."""
    recall_number = record.get("recall_number")
    if not isinstance(recall_number, str) or not recall_number.strip():
        raise ValueError("record has no recall_number")
    return {
        "recall_number": recall_number.strip(),
        "status": str(record.get("status", "unknown")),
        "classification": str(record.get("classification", "unknown")),
        "product_description": str(record.get("product_description", "")),
        "reason_for_recall": str(record.get("reason_for_recall", "")),
        "report_date": record.get("report_date"),
        "source": ENDPOINT,
    }


def fetch(limit: int, api_key: str = "") -> tuple[list[dict[str, Any]], bytes, str]:
    params = {"limit": str(limit), "sort": "report_date:desc"}
    if api_key:
        params["api_key"] = api_key
    url = ENDPOINT + "?" + urllib.parse.urlencode(params)
    req = urllib.request.Request(url, headers={"User-Agent": "toxindex-evidence-challenge/1.0"})
    with urllib.request.urlopen(req, timeout=30) as response:
        raw = response.read()
    payload = json.loads(raw)
    results = payload.get("results")
    if not isinstance(results, list):
        raise ValueError("openFDA response has no results array")
    return [normalize(item) for item in results], raw, url


def write(records: list[dict[str, Any]], raw: bytes, url: str, output: pathlib.Path) -> None:
    output.parent.mkdir(parents=True, exist_ok=True)
    with output.open("w", encoding="utf-8") as handle:
        for record in records:
            handle.write(json.dumps(record, sort_keys=True) + "\n")
    provenance = {
        "retrieved_at": dt.datetime.now(dt.timezone.utc).isoformat(),
        "request_url": url.replace(os.getenv("OPENFDA_API_KEY", ""), "REDACTED")
        if os.getenv("OPENFDA_API_KEY")
        else url,
        "raw_sha256": hashlib.sha256(raw).hexdigest(),
        "records": len(records),
    }
    output.with_suffix(output.suffix + ".provenance.json").write_text(
        json.dumps(provenance, indent=2) + "\n", encoding="utf-8"
    )


def main() -> None:
    parser = argparse.ArgumentParser()
    parser.add_argument("--limit", type=int, default=10, choices=range(1, 101))
    parser.add_argument("--output", type=pathlib.Path, default=pathlib.Path("out/recalls.ndjson"))
    args = parser.parse_args()
    records, raw, url = fetch(args.limit, os.getenv("OPENFDA_API_KEY", ""))
    write(records, raw, url, args.output)
    print(f"wrote {len(records)} records to {args.output}")


if __name__ == "__main__":
    main()

