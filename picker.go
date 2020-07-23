package april

import (
	"github.com/barbosaigor/april/util"
	"github.com/barbosaigor/graphdeppicker"
	"github.com/barbosaigor/graphll"
)

// Pick picks random nodes described in a yaml file as a slice of byte
func Pick(data []byte, quantity uint32) ([]string, error) {
	conf, err := ReadConf(data)
	if err != nil {
		return nil, err
	}
	return PickFromConf(conf, quantity)
}

// PickFromConf picks random nodes from ConfData datastructure
func PickFromConf(conf *ConfData, quantity uint32) ([]string, error) {
	depGraph := graphll.New()
	for sname, sdata := range conf.Services {
		depGraph.Add(sname, sdata.Weight, sdata.Dependencies)
	}
	return graphdeppicker.Pick(depGraph, quantity)
}

// PickFromYaml picks random nodes as PickRandDeps, but it read a yaml file
func PickFromYaml(filePath string, quantity uint32) ([]string, error) {
	fdata, err := util.ReadFile(filePath)
	if err != nil {
		return nil, err
	}
	return Pick(fdata, quantity)
}
