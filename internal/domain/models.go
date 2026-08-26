package domain

import "time"

type Record struct {
	ID        string
	StoreID   string
	AccountID string
	Status    string
	CreatedAt time.Time
	UpdatedAt time.Time
	Notes     string
}
type Profile struct {
	ID      string
	StoreID string
	Name    string
	Phone   string
	Email   string
	Active  bool
	Tags    []string
}
type Event struct {
	ID       string
	RecordID string
	Kind     string
	Actor    string
	At       time.Time
	Payload  string
}
type Audit struct {
	ID       string
	Entity   string
	EntityID string
	Action   string
	Actor    string
	At       time.Time
	Detail   string
}
type Account struct {
	ID        string
	StoreID   string
	Number    string
	Owner     string
	Balance   int64
	State     string
	UpdatedAt time.Time
}

func NewRecord(id, store, account string) Record {
	now := time.Now().UTC()
	return Record{ID: id, StoreID: store, AccountID: account, Status: "received", CreatedAt: now, UpdatedAt: now}
}
func (r Record) IsClosed() bool { return r.Status == "archived" || r.Status == "rejected" }
func (p Profile) Valid() bool   { return p.ID != "" && p.StoreID != "" && p.Name != "" && p.Active }
func (a Account) Usable() bool {
	return a.ID != "" && a.StoreID != "" && a.Number != "" && a.State == "active"
}
