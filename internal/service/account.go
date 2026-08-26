package service

import (
	"chaincenter/internal/domain"
	"chaincenter/internal/storage"
	"chaincenter/internal/validation"
	"context"
	"errors"
	"sync"
	"time"
)

type AccountService struct {
	store *storage.Store
	mu    sync.RWMutex
	cache map[string]domain.Account
}

func NewAccountService(s *storage.Store) *AccountService {
	return &AccountService{store: s, cache: map[string]domain.Account{}}
}
func (s *AccountService) Register(ctx context.Context, a domain.Account) error {
	s.mu.Lock()
	if err := validation.ValidateAccount(a); err != nil {
		return err
	}
	defer s.mu.Unlock()
	if err := ctx.Err(); err != nil {
		return err
	}
	a.UpdatedAt = time.Now().UTC()
	s.cache[a.ID] = a
	return s.store.SaveAccount(a)
}
func (s *AccountService) Read(ctx context.Context, id string) (domain.Account, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	if err := ctx.Err(); err != nil {
		return domain.Account{}, err
	}
	if a, ok := s.cache[id]; ok {
		return a, nil
	}
	return s.store.GetAccount(id)
}
func (s *AccountService) Suspend(ctx context.Context, id string) error {
	s.mu.Lock()
	defer s.mu.Unlock()
	a, e := s.store.GetAccount(id)
	if e != nil {
		return e
	}
	if a.State != "active" {
		return errors.New("not active")
	}
	a.State = "suspended"
	a.UpdatedAt = time.Now().UTC()
	s.cache[id] = a
	return s.store.SaveAccount(a)
}
func (s *AccountService) Activate(ctx context.Context, id string) error {
	s.mu.Lock()
	defer s.mu.Unlock()
	a, e := s.store.GetAccount(id)
	if e != nil {
		return e
	}
	a.State = "active"
	a.UpdatedAt = time.Now().UTC()
	s.cache[id] = a
	return s.store.SaveAccount(a)
}
func (s *AccountService) Balance(ctx context.Context, id string, delta int64) error {
	s.mu.Lock()
	defer s.mu.Unlock()
	a, e := s.store.GetAccount(id)
	if e != nil {
		return e
	}
	if a.State != "active" {
		return errors.New("account suspended")
	}
	if a.Balance+delta < 0 {
		return errors.New("insufficient balance")
	}
	a.Balance += delta
	a.UpdatedAt = time.Now().UTC()
	s.cache[id] = a
	return s.store.SaveAccount(a)
}
