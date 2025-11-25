package kms

import "encoding/json"

// InspectionInfo keeps the SOAP-wrapped inspection payload.
type InspectionInfo struct {
	Raw json.RawMessage `json:"raw"`
}
