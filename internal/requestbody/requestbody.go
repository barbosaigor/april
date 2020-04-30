package requestbody

// ServiceBodyJSON describe a service as JSON structure
type ServiceBodyJSON struct {
	Name     string `json:"name"`
	Selector string `json:"selector"`
}

// ShutdownBodyJSON describe a list of services to shutdown
type ShutdownBodyJSON struct {
	Services []ServiceBodyJSON `json:"services"`
}

// ResponseMessage describe an response message for a JSON structure
type ResponseMessage struct {
	Message string `json:"message"`
}
