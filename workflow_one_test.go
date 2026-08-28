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

func TestWorkflowOne(t *testing.T) {
	s, _ := storage.Open(filepath.Join(t.TempDir(), "x"))
	defer s.Close()
	a := service.NewAccountService(s)
	r := service.NewRecordService(s, a)
	e := workflow.NewEngine(r)
	x := domain.NewRecord("one", "s", "a")
	if err := e.Intake(context.Background(), x, "u"); err != nil {
		t.Fatal(err)
	}
}
