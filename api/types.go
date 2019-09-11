package api

import "github.com/Scalingo/link/models"

const (
	Activated = "ACTIVATED"
	Standby   = "STANDBY"
	Failing   = "FAILING"
)

const (
	TCPHealthCheck = models.TCPHealthCheck
)

type IP struct {
	models.IP
	Status string `json:"status,omitempty"`
}

type UpdateIPOpts struct {
	Checks []models.Healthcheck
}
