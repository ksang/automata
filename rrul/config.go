package rrul

// Config defines netserver related configurations
type Config struct {
	Host        string
	ControlPort uint
	DataPort    uint
	Seconds     uint
}
