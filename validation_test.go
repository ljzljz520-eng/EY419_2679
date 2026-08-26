package main

import (
	"chaincenter/internal/domain"
	"chaincenter/internal/validation"
	"testing"
)

func TestValidationRules(t *testing.T) {
	if validation.ValidateRecord(domain.Record{}) == nil {
		t.Fatal("expected error")
	}
	if validation.NormalizeName(" a  b ") != "a b" {
		t.Fatal("normalize")
	}
}
