package web

const (
	// HealthzLivenessPath is the path for the liveness probe
	HealthzLivenessPath = "/healthz/liveness"
	// HealthzReadinessPath is the path for the readiness probe
	HealthzReadinessPath = "/healthz/readiness"
)
