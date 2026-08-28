package main

import (
	"chaincenter/internal/domain"
	"chaincenter/internal/reporting"
	"testing"
)

func TestReporting(t *testing.T) {
	r := domain.NewRecord("r", "s", "a")
	if reporting.StatusLabel(r.Status) != "Received" {
		t.Fatal("label")
	}
	if reporting.CompletionPercent(r.Status) != 10 {
		t.Fatal("percent")
	}
}
