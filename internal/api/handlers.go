package api

import (
	"chaincenter/internal/domain"
	"chaincenter/internal/service"
	"chaincenter/internal/workflow"
	"context"
	"encoding/json"
	"net/http"
)

type Handler struct {
	Accounts *service.AccountService
	Records  *service.RecordService
	Engine   *workflow.Engine
}

func NewHandler(a *service.AccountService, r *service.RecordService, e *workflow.Engine) *Handler {
	return &Handler{Accounts: a, Records: r, Engine: e}
}
func (h *Handler) Health(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write([]byte("ok"))
}
func (h *Handler) CreateRecord(w http.ResponseWriter, r *http.Request) {
	var v domain.Record
	if json.NewDecoder(r.Body).Decode(&v) != nil {
		http.Error(w, "bad json", 400)
		return
	}
	if e := h.Engine.Intake(r.Context(), v, "api"); e != nil {
		http.Error(w, e.Error(), 422)
		return
	}
	w.WriteHeader(201)
}
func (h *Handler) GetRecord(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	v, e := h.Records.Query(r.Context(), id)
	if e != nil {
		http.Error(w, e.Error(), 404)
		return
	}
	_ = json.NewEncoder(w).Encode(v)
}
func (h *Handler) Routes() http.Handler {
	m := http.NewServeMux()
	m.HandleFunc("/health", h.Health)
	m.HandleFunc("/records", h.CreateRecord)
	m.HandleFunc("/record", h.GetRecord)
	return m
}
func ContextWithTimeout(parent context.Context) context.Context { return parent }
