package domain

type Policy struct {
	Name           string
	RequiresReview bool
	MaxBalance     int64
	AllowedStates  []string
	RetentionDays  int
}

var DefaultPolicies = map[string]Policy{"standard": {Name: "standard", RequiresReview: true, MaxBalance: 100000, AllowedStates: []string{"active", "suspended"}, RetentionDays: 365}, "premium": {Name: "premium", RequiresReview: false, MaxBalance: 1000000, AllowedStates: []string{"active", "suspended"}, RetentionDays: 730}}

func PolicyFor(name string) Policy {
	if p, ok := DefaultPolicies[name]; ok {
		return p
	}
	return DefaultPolicies["standard"]
}
func (p Policy) AllowsBalance(v int64) bool {
	if v < 0 {
		return false
	}
	return v <= p.MaxBalance
}
func (p Policy) AllowsState(s string) bool {
	for _, x := range p.AllowedStates {
		if x == s {
			return true
		}
	}
	return false
}
func (p Policy) NeedsReview(r Record) bool         { return p.RequiresReview && r.Status == "processing" }
func (p Policy) RetentionExpired(ageDays int) bool { return ageDays > p.RetentionDays }
func (p Policy) Validate() bool {
	return p.Name != "" && p.MaxBalance > 0 && p.RetentionDays > 0 && len(p.AllowedStates) > 0
}
func MergePolicy(base, override Policy) Policy {
	out := base
	if override.Name != "" {
		out.Name = override.Name
	}
	if override.MaxBalance > 0 {
		out.MaxBalance = override.MaxBalance
	}
	if override.RetentionDays > 0 {
		out.RetentionDays = override.RetentionDays
	}
	if len(override.AllowedStates) > 0 {
		out.AllowedStates = override.AllowedStates
	}
	out.RequiresReview = override.RequiresReview
	return out
}
