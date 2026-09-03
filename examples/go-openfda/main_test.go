package main

import "testing"

func TestNormalizePreservesSourceIdentity(t *testing.T) {
	got, err := normalize(map[string]any{"recall_number": " D-123 ", "status": "Ongoing"})
	if err != nil {
		t.Fatal(err)
	}
	if got.RecallNumber != "D-123" || got.Status != "Ongoing" || got.Source != endpoint {
		t.Fatalf("unexpected normalized record: %+v", got)
	}
}

func TestNormalizeRejectsMissingIdentity(t *testing.T) {
	if _, err := normalize(map[string]any{"status": "Ongoing"}); err == nil {
		t.Fatal("expected missing recall_number to fail")
	}
}
