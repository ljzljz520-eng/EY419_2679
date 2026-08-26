package service

import (
	"chaincenter/internal/domain"
	"chaincenter/internal/storage"
	"chaincenter/internal/validation"
	"context"
	"fmt"
	"time"
)

type RecordService struct {
	store    *storage.Store
	accounts *AccountService
}

func NewRecordService(s *storage.Store, a *AccountService) *RecordService {
	return &RecordService{store: s, accounts: a}
}
func (s *RecordService) Receive(ctx context.Context, r domain.Record) error {
	if e := ctx.Err(); e != nil {
		return e
	}
	if e := validation.ValidateRecord(r); e != nil {
		return e
	}
	return s.store.SaveRecord(r)
}
func (s *RecordService) Advance(ctx context.Context, id, to, actor string) error {
	r, e := s.store.GetRecord(id)
	if e != nil {
		return e
	}
	if e = validation.ValidateTransition(r.Status, to); e != nil {
		return e
	}
	if e = ctx.Err(); e != nil {
		return e
	}
	r.Status = to
	r.UpdatedAt = time.Now().UTC()
	if e = s.store.SaveRecord(r); e != nil {
		return e
	}
	return s.store.SaveEvent(domain.Event{ID: fmt.Sprintf("%s-%d", id, time.Now().UnixNano()), RecordID: id, Kind: "status", Actor: actor, At: time.Now().UTC(), Payload: to})
}
func (s *RecordService) Archive(ctx context.Context, id, actor string) error {
	return s.Advance(ctx, id, "archived", actor)
}
func (s *RecordService) Query(ctx context.Context, id string) (domain.Record, error) {
	if e := ctx.Err(); e != nil {
		return domain.Record{}, e
	}
	return s.store.GetRecord(id)
}
func (s *RecordService) AddNote(ctx context.Context, id, note string) error {
	r, e := s.store.GetRecord(id)
	if e != nil {
		return e
	}
	r.Notes = note
	r.UpdatedAt = time.Now().UTC()
	return s.store.SaveRecord(r)
}
