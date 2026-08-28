package workflow

import (
	"chaincenter/internal/domain"
	"chaincenter/internal/service"
	"context"
	"fmt"
)

type Engine struct{ records *service.RecordService }

func NewEngine(r *service.RecordService) *Engine { return &Engine{records: r} }
func (e *Engine) Intake(ctx context.Context, r domain.Record, actor string) error {
	if err := e.records.Receive(ctx, r); err != nil {
		return err
	}
	if err := e.records.Advance(ctx, r.ID, "validated", actor); err != nil {
		return err
	}
	return e.records.Advance(ctx, r.ID, "processing", actor)
}
func (e *Engine) Review(ctx context.Context, id, actor string) error {
	if err := e.records.Advance(ctx, id, "approved", actor); err != nil {
		return err
	}
	return nil
}
func (e *Engine) Archive(ctx context.Context, id, actor string) error {
	if err := e.records.Archive(ctx, id, actor); err != nil {
		return fmt.Errorf("archive: %w", err)
	}
	return nil
}
func (e *Engine) Track(ctx context.Context, id string) (domain.Record, error) {
	return e.records.Query(ctx, id)
}
