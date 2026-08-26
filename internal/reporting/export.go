package reporting

import (
	"chaincenter/internal/domain"
	"encoding/json"
	"fmt"
)

func ToJSON(r domain.Record) (string, error) {
	b, e := json.MarshalIndent(r, "", "  ")
	return string(b), e
}
func ToCSV(r domain.Record) string {
	return fmt.Sprintf("%s,%s,%s,%s", r.ID, r.StoreID, r.AccountID, r.Status)
}
func StatusLabel(s string) string {
	switch s {
	case "received":
		return "Received"
	case "validated":
		return "Validated"
	case "processing":
		return "Processing"
	case "approved":
		return "Approved"
	case "archived":
		return "Archived"
	case "rejected":
		return "Rejected"
	}
	return "Unknown"
}
func CompletionPercent(s string) int {
	switch s {
	case "received":
		return 10
	case "validated":
		return 35
	case "processing":
		return 60
	case "approved":
		return 85
	case "archived":
		return 100
	}
	return 0
}
