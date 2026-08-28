package main

import (
	"chaincenter/internal/domain"
	"chaincenter/internal/storage"
	"path/filepath"
	"testing"
)

func TestPersistenceSurvivesReopen(t *testing.T) {
	p := filepath.Join(t.TempDir(), "data.db")
	s, e := storage.Open(p)
	if e != nil {
		t.Fatal(e)
	}
	r := domain.NewRecord("r1", "store", "a1")
	if e = s.SaveRecord(r); e != nil {
		t.Fatal(e)
	}
	s.Close()
	s, e = storage.Open(p)
	if e != nil {
		t.Fatal(e)
	}
	defer s.Close()
	got, e := s.GetRecord("r1")
	if e != nil || got.ID != "r1" {
		t.Fatalf("%v %+v", e, got)
	}
}
