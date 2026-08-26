package main

import (
	"chaincenter/internal/domain"
	"chaincenter/internal/service"
	"chaincenter/internal/storage"
	"context"
	"path/filepath"
	"testing"
	"time"
)

func TestBusinessChain28(t *testing.T) {
	s, _ := storage.Open(filepath.Join(t.TempDir(), "x"))
	defer s.Close()
	a := service.NewAccountService(s)
	bad := domain.Account{ID: "bad", StoreID: "s", Number: "1", State: "active"}
	_ = a.Register(context.Background(), bad)
	done := make(chan struct{})
	go func() { _, _ = a.Read(context.Background(), "missing"); close(done) }()
	select {
	case <-done:
	case <-time.After(100 * time.Millisecond):
		t.Fatal("read blocked after validation")
	}
}
