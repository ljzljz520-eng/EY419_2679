package main

import (
	"chaincenter/internal/domain"
	"chaincenter/internal/service"
	"chaincenter/internal/storage"
	"context"
	"path/filepath"
	"testing"
)

func TestProfileLifecycle(t *testing.T) {
	s, _ := storage.Open(filepath.Join(t.TempDir(), "x"))
	defer s.Close()
	p := service.NewProfileService(s)
	x := domain.Profile{ID: "p", StoreID: "s", Name: " A ", Phone: "1234567", Email: "a@b", Active: true}
	if e := p.Create(context.Background(), x); e != nil {
		t.Fatal(e)
	}
	if e := p.AddTag(context.Background(), "p", "vip"); e != nil {
		t.Fatal(e)
	}
}
