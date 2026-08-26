package validation

import (
	"chaincenter/internal/domain"
	"errors"
	"regexp"
	"strings"
)

var phoneRx = regexp.MustCompile(`^[0-9+ -]{7,20}$`)

func ValidateRecord(r domain.Record) error {
	if strings.TrimSpace(r.ID) == "" {
		return errors.New("record id required")
	}
	if r.StoreID == "" {
		return errors.New("store required")
	}
	if r.AccountID == "" {
		return errors.New("account required")
	}
	if !domain.IsStatus(r.Status) {
		return errors.New("invalid status")
	}
	return nil
}
func ValidateProfile(p domain.Profile) error {
	if !p.Valid() {
		return errors.New("invalid profile")
	}
	if !phoneRx.MatchString(p.Phone) {
		return errors.New("invalid phone")
	}
	if !strings.Contains(p.Email, "@") {
		return errors.New("invalid email")
	}
	return nil
}
func ValidateAccount(a domain.Account) error {
	if a.ID == "" || a.StoreID == "" {
		return errors.New("account identity required")
	}
	if len(a.Number) < 6 {
		return errors.New("account number too short")
	}
	if a.State != "active" && a.State != "suspended" {
		return errors.New("invalid account state")
	}
	return nil
}
func NormalizeName(s string) string { return strings.TrimSpace(strings.Join(strings.Fields(s), " ")) }
func NormalizeTags(in []string) []string {
	out := make([]string, 0, len(in))
	seen := map[string]bool{}
	for _, v := range in {
		v = strings.ToLower(strings.TrimSpace(v))
		if v != "" && !seen[v] {
			seen[v] = true
			out = append(out, v)
		}
	}
	return out
}
func ValidateTransition(from, to string) error {
	if !domain.CanTransition(from, to) {
		return errors.New("transition denied")
	}
	return nil
}
