package audit

import (
	"chaincenter/internal/domain"
	"chaincenter/internal/storage"
	"fmt"
	"time"
)

type Logger struct{ store *storage.Store }

func New(s *storage.Store) *Logger { return &Logger{store: s} }
func (l *Logger) Log(entity, id, action, actor, detail string) error {
	return l.store.SaveAudit(domain.Audit{ID: fmt.Sprintf("%s-%d", entity, time.Now().UnixNano()), Entity: entity, EntityID: id, Action: action, Actor: actor, At: time.Now().UTC(), Detail: detail})
}
func (l *Logger) RecordCreated(id, actor string) error {
	return l.Log("record", id, "created", actor, "intake")
}
func (l *Logger) RecordArchived(id, actor string) error {
	return l.Log("record", id, "archived", actor, "completed")
}
