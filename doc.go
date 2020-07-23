// Package april is a chaos testing tool.
//
// April proposes to improve resilience in microservices architectures. It
// does chaos testing by randomly shutting down services, taking into account their importance.
// April is similar to other tools, such as Chaos Monkey, Pumba, and Gremlin.
// But it differs by the selection algorithm, which considers services importances (weights).
package april
