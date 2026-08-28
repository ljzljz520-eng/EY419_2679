package main

import (
	"chaincenter/internal/domain"
	"chaincenter/internal/service"
	"chaincenter/internal/storage"
	"chaincenter/internal/workflow"
	"context"
	"path/filepath"
	"testing"
)

func TestWorkflowTwo(t *testing.T) {
	s, _ := storage.Open(filepath.Join(t.TempDir(), "x"))
	defer s.Close()
	a := service.NewAccountService(s)
	r := service.NewRecordService(s, a)
	e := workflow.NewEngine(r)
	x := domain.NewRecord("two", "s", "a")
	_ = e.Intake(context.Background(), x, "u")
	if err := e.Review(context.Background(), x.ID, "reviewer"); err != nil {
		t.Fatal(err)
	}
	if err := e.Archive(context.Background(), x.ID, "archiver"); err != nil {
		t.Fatal(err)
	}
}
