package models

type IP struct {
	ID      string        `json:"id"`
	IP      string        `json:"ip"`
	LeaseID int64         `json:"lease_id,omitempty"`
	Status  string        `json:"status,omitempty"`
	Checks  []Healthcheck `json:"checks,omitempty"`
}
