package main

import (
	"chaincenter/internal/api"
	"chaincenter/internal/service"
	"chaincenter/internal/storage"
	"chaincenter/internal/workflow"
	"net/http/httptest"
	"path/filepath"
	"testing"
)

func TestHealthEndpoint(t *testing.T) {
	s, _ := storage.Open(filepath.Join(t.TempDir(), "x"))
	defer s.Close()
	a := service.NewAccountService(s)
	r := service.NewRecordService(s, a)
	h := api.NewHandler(a, r, workflow.NewEngine(r))
	w := httptest.NewRecorder()
	h.Health(w, httptest.NewRequest("GET", "/health", nil))
	if w.Code != 200 {
		t.Fatal(w.Code)
	}
}
