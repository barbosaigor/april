package april

import (
	"gopkg.in/yaml.v2"
)

// ConfService stores service data
type ConfService map[string]struct {
	Weight       uint32   `yaml:"weight"`
	Dependencies []string `yaml:"dependencies"`
	Selector     string   `yaml:"selector"`
}

// ConfData stores the configuration data
type ConfData struct {
	Version  int32       `yaml:"version"`
	Services ConfService `yaml:"services"`
}

// ReadConf reads bytes and convert to an service
// data structure
func ReadConf(conf []byte) (*ConfData, error) {
	var servs ConfData
	if err := yaml.Unmarshal(conf, &servs); err != nil {
		return nil, err
	}
	return &servs, nil
}
