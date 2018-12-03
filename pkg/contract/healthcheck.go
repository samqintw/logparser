package contract

type HealthCheckRequest struct {
	Name string
	MailAddress []string
	Log []byte
}

type HealthCheckResponse struct {
	Message string
}
