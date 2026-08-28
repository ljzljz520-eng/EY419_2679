package main

import (
	"chaincenter/internal/domain"
	"chaincenter/internal/service"
	"chaincenter/internal/storage"
	"context"
	"path/filepath"
	"testing"
)

func TestAccountLifecycle(t *testing.T) {
	s, _ := storage.Open(filepath.Join(t.TempDir(), "x"))
	defer s.Close()
	a := service.NewAccountService(s)
	x := domain.Account{ID: "a", StoreID: "s", Number: "123456", State: "active"}
	if e := a.Register(context.Background(), x); e != nil {
		t.Fatal(e)
	}
	if e := a.Suspend(context.Background(), x.ID); e != nil {
		t.Fatal(e)
	}
	if e := a.Activate(context.Background(), x.ID); e != nil {
		t.Fatal(e)
	}
}
