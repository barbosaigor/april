package destroyer

type Status uint8
// Status allowed values
const (
	Up = Status(iota)
	Down
	Pending
	Any
)

type Instance struct {
	Name string
	Sts Status
}
