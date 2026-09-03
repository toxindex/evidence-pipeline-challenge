import unittest

from src.ingest import normalize


class NormalizeTest(unittest.TestCase):
    def test_preserves_source_identity(self):
        row = normalize({"recall_number": " D-123 ", "status": "Ongoing"})
        self.assertEqual(row["recall_number"], "D-123")
        self.assertEqual(row["status"], "Ongoing")
        self.assertIn("fda.gov", row["source"])

    def test_rejects_record_without_identity(self):
        with self.assertRaisesRegex(ValueError, "recall_number"):
            normalize({"status": "Ongoing"})


if __name__ == "__main__":
    unittest.main()

