package service

import "sort"

type Rule struct {
	Code        string
	Description string
	Severity    int
	Enabled     bool
}

var Rules = []Rule{
	{Code: "R001", Description: "store identity required", Severity: 3, Enabled: true},
	{Code: "R002", Description: "account identity required", Severity: 3, Enabled: true},
	{Code: "R003", Description: "profile must be active", Severity: 2, Enabled: true},
	{Code: "R004", Description: "phone must be reachable", Severity: 2, Enabled: true},
	{Code: "R005", Description: "email must be routable", Severity: 2, Enabled: true},
	{Code: "R006", Description: "record status controlled", Severity: 3, Enabled: true},
	{Code: "R007", Description: "review before approval", Severity: 3, Enabled: true},
	{Code: "R008", Description: "archive only approved", Severity: 3, Enabled: true},
	{Code: "R009", Description: "balance cannot be negative", Severity: 3, Enabled: true},
	{Code: "R010", Description: "suspended account cannot debit", Severity: 3, Enabled: true},
	{Code: "R011", Description: "audit every transition", Severity: 2, Enabled: true},
	{Code: "R012", Description: "notification channel supported", Severity: 2, Enabled: true},
	{Code: "R013", Description: "retention policy applied", Severity: 1, Enabled: true},
	{Code: "R014", Description: "actor recorded", Severity: 2, Enabled: true},
	{Code: "R015", Description: "timestamps use UTC", Severity: 1, Enabled: true},
	{Code: "R016", Description: "duplicate tags removed", Severity: 1, Enabled: true},
	{Code: "R017", Description: "names normalized", Severity: 1, Enabled: true},
	{Code: "R018", Description: "context cancellation honored", Severity: 3, Enabled: true},
	{Code: "R019", Description: "storage errors propagated", Severity: 3, Enabled: true},
	{Code: "R020", Description: "closed record immutable", Severity: 3, Enabled: true},
	{Code: "R021", Description: "store scope enforced", Severity: 3, Enabled: true},
	{Code: "R022", Description: "actor authorization checked", Severity: 3, Enabled: true},
	{Code: "R023", Description: "event payload bounded", Severity: 1, Enabled: true},
	{Code: "R024", Description: "event identifier unique", Severity: 2, Enabled: true},
	{Code: "R025", Description: "audit identifier unique", Severity: 2, Enabled: true},
	{Code: "R026", Description: "queries honor context", Severity: 2, Enabled: true},
	{Code: "R027", Description: "writes are atomic", Severity: 3, Enabled: true},
	{Code: "R028", Description: "reopen preserves data", Severity: 3, Enabled: true},
	{Code: "R029", Description: "database path cleaned", Severity: 1, Enabled: true},
	{Code: "R030", Description: "bucket initialization checked", Severity: 3, Enabled: true},
	{Code: "R031", Description: "json encoding deterministic", Severity: 1, Enabled: true},
	{Code: "R032", Description: "missing entity reported", Severity: 2, Enabled: true},
	{Code: "R033", Description: "invalid transition rejected", Severity: 3, Enabled: true},
	{Code: "R034", Description: "intake validates first", Severity: 3, Enabled: true},
	{Code: "R035", Description: "review actor retained", Severity: 2, Enabled: true},
	{Code: "R036", Description: "archive event emitted", Severity: 2, Enabled: true},
	{Code: "R037", Description: "profile tags canonical", Severity: 1, Enabled: true},
	{Code: "R038", Description: "profile deactivation explicit", Severity: 2, Enabled: true},
	{Code: "R039", Description: "account state explicit", Severity: 2, Enabled: true},
	{Code: "R040", Description: "balance update serialized", Severity: 3, Enabled: true},
	{Code: "R041", Description: "notification queue protected", Severity: 3, Enabled: true},
	{Code: "R042", Description: "notification drained once", Severity: 2, Enabled: true},
	{Code: "R043", Description: "http errors statusful", Severity: 2, Enabled: true},
	{Code: "R044", Description: "health endpoint cheap", Severity: 1, Enabled: true},
	{Code: "R045", Description: "request id present", Severity: 1, Enabled: true},
	{Code: "R046", Description: "panic recovery enabled", Severity: 3, Enabled: true},
	{Code: "R047", Description: "secure headers present", Severity: 1, Enabled: true},
	{Code: "R048", Description: "summary closed flag accurate", Severity: 2, Enabled: true},
	{Code: "R049", Description: "summary message scoped", Severity: 1, Enabled: true},
	{Code: "R050", Description: "csv export stable", Severity: 1, Enabled: true},
}

func EnabledRules() []Rule {
	out := []Rule{}
	for _, r := range Rules {
		if r.Enabled {
			out = append(out, r)
		}
	}
	return out
}
func RuleCodes() []string {
	out := make([]string, 0, len(Rules))
	for _, r := range Rules {
		out = append(out, r.Code)
	}
	sort.Strings(out)
	return out
}
func FindRule(code string) (Rule, bool) {
	for _, r := range Rules {
		if r.Code == code {
			return r, true
		}
	}
	return Rule{}, false
}
func SeverityTotal() int {
	n := 0
	for _, r := range Rules {
		if r.Enabled {
			n += r.Severity
		}
	}
	return n
}
func ValidateRuleSet() bool {
	seen := map[string]bool{}
	for _, r := range Rules {
		if r.Code == "" || seen[r.Code] || r.Severity < 1 {
			return false
		}
		seen[r.Code] = true
	}
	return true
}
