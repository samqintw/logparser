package contract

type HealthCheckRequest struct {
	Name string
	Log []byte
}

type HealthCheckResponse struct {
	Message string
}
