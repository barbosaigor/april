package destroyer

// LifeCycle contains routines that execute in the chaos server lifecycle
type LifeCycle interface {
	// OnStart runs before the host is served.
	// Could be used as configuration routine for before stating the server
	OnStart()
}

// ChaosServer is an interface for chaos server implementation
type ChaosServer interface {
	// Shutdown turns down instances.
	// For Chaos Server implementation especific
	Shutdown(svc string) error
	// ListInstances lists all instances that are in any state
	ListInstances(status Status) ([]Instance, error)
	LifeCycle
}
