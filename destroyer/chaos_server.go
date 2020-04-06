package destroyer

type ChaosServer interface {
	// Shutdown turns down instances.
	// For Chaos Server implementation especific
	Shutdown(svc string) error
	// ListInstances lists all instances that are in any state
	ListInstances(status Status) ([]Instance, error)
}
