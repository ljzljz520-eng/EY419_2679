package storage

import (
	"chaincenter/internal/domain"
	"encoding/json"
	"fmt"
	"go.etcd.io/bbolt"
	"path/filepath"
	"sync"
)

var buckets = []string{"records", "profiles", "events", "audits", "accounts"}

type Store struct {
	db   *bbolt.DB
	mu   sync.RWMutex
	path string
}

func Open(path string) (*Store, error) {
	db, e := bbolt.Open(filepath.Clean(path), 0600, nil)
	if e != nil {
		return nil, e
	}
	s := &Store{db: db, path: path}
	e = db.Update(func(tx *bbolt.Tx) error {
		for _, b := range buckets {
			if _, x := tx.CreateBucketIfNotExists([]byte(b)); x != nil {
				return x
			}
		}
		return nil
	})
	if e != nil {
		db.Close()
		return nil, e
	}
	return s, nil
}
func (s *Store) Close() error {
	s.mu.Lock()
	defer s.mu.Unlock()
	if s.db == nil {
		return nil
	}
	e := s.db.Close()
	s.db = nil
	return e
}
func enc(v any) ([]byte, error) { return json.Marshal(v) }
func put(db *bbolt.DB, bucket, key string, v any) error {
	data, e := enc(v)
	if e != nil {
		return e
	}
	return db.Update(func(tx *bbolt.Tx) error { return tx.Bucket([]byte(bucket)).Put([]byte(key), data) })
}
func get(db *bbolt.DB, bucket, key string, out any) error {
	return db.View(func(tx *bbolt.Tx) error {
		v := tx.Bucket([]byte(bucket)).Get([]byte(key))
		if v == nil {
			return fmt.Errorf("not found")
		}
		return json.Unmarshal(v, out)
	})
}
func (s *Store) SaveRecord(v domain.Record) error {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return put(s.db, "records", v.ID, v)
}
func (s *Store) GetRecord(id string) (domain.Record, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	var v domain.Record
	e := get(s.db, "records", id, &v)
	return v, e
}
func (s *Store) SaveProfile(v domain.Profile) error {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return put(s.db, "profiles", v.ID, v)
}
func (s *Store) GetProfile(id string) (domain.Profile, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	var v domain.Profile
	e := get(s.db, "profiles", id, &v)
	return v, e
}
func (s *Store) SaveEvent(v domain.Event) error {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return put(s.db, "events", v.ID, v)
}
func (s *Store) SaveAudit(v domain.Audit) error {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return put(s.db, "audits", v.ID, v)
}
func (s *Store) SaveAccount(v domain.Account) error {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return put(s.db, "accounts", v.ID, v)
}
func (s *Store) GetAccount(id string) (domain.Account, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	var v domain.Account
	e := get(s.db, "accounts", id, &v)
	return v, e
}
