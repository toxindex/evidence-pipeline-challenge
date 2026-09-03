import unittest

from src.ingest import Platform, build_bundle


class FakePlatform:
    def bricks(self): return {"items": [{"slug": "tox21"}]}
    def search(self, query): return {"hits": [{"brick": "tox21", "snippet": query}]}
    def prediction_tools(self): return {"items": [{"name": "admet-ai"}]}
    def predict(self, tool, smiles): return {"tool": tool, "input": smiles, "score": 0.4}


class EvidenceBundleTest(unittest.TestCase):
    def test_keeps_source_and_prediction_provenance(self):
        bundle = build_bundle(FakePlatform(), "acetaminophen liver", "admet-ai", "CCO")
        self.assertEqual(bundle["sources"]["search"]["hits"][0]["brick"], "tox21")
        self.assertEqual(bundle["prediction"]["requested_tool"], "admet-ai")
        self.assertEqual(bundle["prediction"]["catalog_at_run"]["items"][0]["name"], "admet-ai")

    def test_requires_platform_key(self):
        with self.assertRaisesRegex(ValueError, "TOXINDEX_API_KEY"):
            Platform("")


if __name__ == "__main__":
    unittest.main()

