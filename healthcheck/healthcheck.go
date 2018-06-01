package healthcheck

// HealthCheck - status
type HealthCheck struct {
	Version string `json:"version"`
	Status  string `json:"status"`
}

// GetHealthCheck  is a temporary function used to learn now to test in go
func (h HealthCheck) GetHealthCheck() HealthCheck {
	return h
}
