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

func TestWorkflowThree(t *testing.T) {
	s, _ := storage.Open(filepath.Join(t.TempDir(), "x"))
	defer s.Close()
	a := service.NewAccountService(s)
	r := service.NewRecordService(s, a)
	e := workflow.NewEngine(r)
	x := domain.NewRecord("three", "s", "a")
	if err := e.Intake(context.Background(), x, "u"); err != nil {
		t.Fatal(err)
	}
	if _, err := e.Track(context.Background(), x.ID); err != nil {
		t.Fatal(err)
	}
}
