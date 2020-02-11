package april

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"path/filepath"
)

type service map[string]struct {
	Weight       uint32   `yaml:"weight"`
	Dependencies []string `yaml:"dependencies"`
}

type services struct {
	ServicesData service `yaml:"services"`
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
func getConf(conf []byte) (*services, error) {
	var servs services
	if err := yaml.Unmarshal(conf, &servs); err != nil {
		return nil, err
	}
	return &servs, nil
}

// getConfFile read a yaml file from system and
// convert to a service configuration data structure
func getConfFile(filename string) (*services, error) {
	fdata, err := readFile(filename)
	if err != nil {
		return nil, err
	}
	servs, err := getConf(fdata)
	if err != nil {
		return nil, err
	}
	return servs, nil
}
