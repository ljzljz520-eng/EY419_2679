package domain

var Statuses = []string{"received", "validated", "processing", "approved", "archived", "rejected"}

func IsStatus(s string) bool {
	for _, v := range Statuses {
		if s == v {
			return true
		}
	}
	return false
}
func CanTransition(from, to string) bool {
	if !IsStatus(from) || !IsStatus(to) {
		return false
	}
	switch from {
	case "received":
		return to == "validated" || to == "rejected"
	case "validated":
		return to == "processing" || to == "rejected"
	case "processing":
		return to == "approved" || to == "rejected"
	case "approved":
		return to == "archived"
	}
	return false
}
