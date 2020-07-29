package chaosserver

// Status represents instance status
type Status uint8

// Status allowed values
const (
	Up = Status(iota)
	Down
	Pending
	Any
)

// Instance holds the an instance data
type Instance struct {
	Name string
	Sts  Status
}
