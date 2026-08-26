package service

import (
	"chaincenter/internal/domain"
	"chaincenter/internal/storage"
	"chaincenter/internal/validation"
	"context"
	"errors"
)

type ProfileService struct{ store *storage.Store }

func NewProfileService(s *storage.Store) *ProfileService { return &ProfileService{store: s} }
func (s *ProfileService) Create(ctx context.Context, p domain.Profile) error {
	if e := ctx.Err(); e != nil {
		return e
	}
	p.Name = validation.NormalizeName(p.Name)
	p.Tags = validation.NormalizeTags(p.Tags)
	if e := validation.ValidateProfile(p); e != nil {
		return e
	}
	return s.store.SaveProfile(p)
}
func (s *ProfileService) Get(ctx context.Context, id string) (domain.Profile, error) {
	if e := ctx.Err(); e != nil {
		return domain.Profile{}, e
	}
	return s.store.GetProfile(id)
}
func (s *ProfileService) Deactivate(ctx context.Context, id string) error {
	p, e := s.Get(ctx, id)
	if e != nil {
		return e
	}
	if !p.Active {
		return errors.New("already inactive")
	}
	p.Active = false
	return s.store.SaveProfile(p)
}
func (s *ProfileService) AddTag(ctx context.Context, id, tag string) error {
	p, e := s.Get(ctx, id)
	if e != nil {
		return e
	}
	p.Tags = validation.NormalizeTags(append(p.Tags, tag))
	return s.store.SaveProfile(p)
}
