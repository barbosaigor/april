package april

import (
	"gopkg.in/yaml.v2"
)

type service map[string]struct {
	Weight       uint32   `yaml:"weight"`
	Dependencies []string `yaml:"dependencies"`
	Selector     string   `yaml:"selector"`
}

type ConfData struct {
	Version  int32   `yaml:"version"`
	Services service `yaml:"services"`
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
