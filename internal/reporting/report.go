package reporting

import (
	"chaincenter/internal/domain"
	"chaincenter/internal/service"
	"context"
	"fmt"
)

type Summary struct {
	RecordID string
	Status   string
	Closed   bool
	Message  string
}

func Build(ctx context.Context, s *service.RecordService, id string) (Summary, error) {
	r, e := s.Query(ctx, id)
	if e != nil {
		return Summary{}, e
	}
	return Summary{RecordID: r.ID, Status: r.Status, Closed: r.IsClosed(), Message: fmt.Sprintf("store %s account %s", r.StoreID, r.AccountID)}, nil
}
func Explain(r domain.Record) string {
	if r.IsClosed() {
		return "record is closed"
	}
	return "record remains active"
}
func Group(records []domain.Record) map[string]int {
	out := map[string]int{}
	for _, r := range records {
		out[r.Status]++
	}
	return out
}
