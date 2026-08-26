package storage

import (
	"errors"
	"go.etcd.io/bbolt"
	"time"
)

func (s *Store) Ready() error {
	s.mu.RLock()
	defer s.mu.RUnlock()
	if s.db == nil {
		return errors.New("store closed")
	}
	return s.db.View(func(*bbolt.Tx) error { return nil })
}
func (s *Store) Path() string     { return s.path }
func (s *Store) Touch() time.Time { return time.Now().UTC() }
