package workflow

import (
	"context"
	"fmt"
	"sync"
	"time"
)

type Notification struct {
	ID       string
	RecordID string
	Channel  string
	Message  string
	SentAt   time.Time
}
type Notifier struct {
	mu    sync.Mutex
	queue []Notification
}

func NewNotifier() *Notifier { return &Notifier{queue: []Notification{}} }
func (n *Notifier) Enqueue(ctx context.Context, id, channel, message string) error {
	if e := ctx.Err(); e != nil {
		return e
	}
	n.mu.Lock()
	defer n.mu.Unlock()
	if channel != "email" && channel != "sms" {
		return fmt.Errorf("unsupported channel")
	}
	n.queue = append(n.queue, Notification{ID: fmt.Sprintf("%s-%d", id, time.Now().UnixNano()), RecordID: id, Channel: channel, Message: message})
	return nil
}
func (n *Notifier) Drain(ctx context.Context) []Notification {
	n.mu.Lock()
	defer n.mu.Unlock()
	out := append([]Notification(nil), n.queue...)
	n.queue = nil
	for i := range out {
		out[i].SentAt = time.Now().UTC()
	}
	return out
}
func (n *Notifier) Pending() int { n.mu.Lock(); defer n.mu.Unlock(); return len(n.queue) }
