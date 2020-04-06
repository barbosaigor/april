package april

import (
	"github.com/barbosaigor/graphdeppicker"
	"github.com/barbosaigor/graphll"
)

// pickRandDepsYml picks random nodes from confData datastructure
func pickRandDeps(conf *confData, quantity uint32) ([]string, error) {
	depGraph := graphll.New()
	for sname, sdata := range conf.Services {
		depGraph.Add(sname, sdata.Weight, sdata.Dependencies)
	}
	return graphdeppicker.Run(depGraph, quantity)
}

// PickRandDeps picks random nodes described in a yaml file as a slice of byte
func PickRandDeps(data []byte, quantity uint32) ([]string, error) {
	conf, err := getConf(data)
	if err != nil {
		return nil, err
	}
	return pickRandDeps(conf, quantity)
}

// PickRandDepsYml picks random nodes as PickRandDeps, but it read a yaml file
func PickRandDepsYml(filePath string, quantity uint32) ([]string, error) {
	fdata, err := readFile(filePath)
	if err != nil {
		return nil, err
	}
	return PickRandDeps(fdata, quantity)
}
