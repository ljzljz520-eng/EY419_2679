package storage

import (
	"chaincenter/internal/domain"
	"encoding/json"
	"go.etcd.io/bbolt"
	"strings"
)

func (s *Store) ListRecords(status string) ([]domain.Record, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	out := []domain.Record{}
	e := s.db.View(func(tx *bbolt.Tx) error {
		b := tx.Bucket([]byte("records"))
		return b.ForEach(func(_, v []byte) error {
			var r domain.Record
			if json.Unmarshal(v, &r) != nil {
				return nil
			}
			if status == "" || r.Status == status {
				out = append(out, r)
			}
			return nil
		})
	})
	return out, e
}
func NormalizeKey(v string) string { return strings.TrimSpace(strings.ToLower(v)) }
func MatchRecord(r domain.Record, store, account string) bool {
	if store != "" && r.StoreID != store {
		return false
	}
	if account != "" && r.AccountID != account {
		return false
	}
	return true
}
func SortRecords(in []domain.Record) []domain.Record {
	out := append([]domain.Record(nil), in...)
	for i := 0; i < len(out); i++ {
		for j := i + 1; j < len(out); j++ {
			if out[j].UpdatedAt.Before(out[i].UpdatedAt) {
				out[i], out[j] = out[j], out[i]
			}
		}
	}
	return out
}
