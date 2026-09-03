"""INTENTIONALLY FLAWED pseudocode-like Temporal starter."""
from temporalio import workflow
import requests


@workflow.defn
class RefreshEvidence:
    @workflow.run
    async def run(self, chemical):
        # Faults: network I/O inside deterministic workflow code, secret read is
        # nondeterministic, no activity/retry boundary, result has no source id.
        return requests.get(
            "https://platform.toxindex.com/v1/kiln/search",
            params={"q": chemical},
        ).json()

