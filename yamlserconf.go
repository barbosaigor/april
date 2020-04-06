package april

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"path/filepath"
)

type service map[string]struct {
	Weight       uint32   `yaml:"weight"`
	Dependencies []string `yaml:"dependencies"`
	Selector     string   `yaml:"selector"`
}

type confData struct {
	Version  int32   `yaml:"version"`
	Services service `yaml:"services"`
}

func readFile(filename string) ([]byte, error) {
	fname, err := filepath.Abs(filename)
	if err != nil {
		return nil, err
	}
	fdata, err := ioutil.ReadFile(fname)
	if err != nil {
		return nil, err
	}
	return fdata, nil
}

// getConf reads bytes and convert to an service
// data structure
func getConf(conf []byte) (*confData, error) {
	var servs confData
	if err := yaml.Unmarshal(conf, &servs); err != nil {
		return nil, err
	}
	return &servs, nil
}
