package april

import (
	"github.com/barbosaigor/graphdeppicker"
	"github.com/barbosaigor/graphll"
)

// pickRandDepsYml picks random nodes from services datastructure
func pickRandDeps(servs *services, quantity uint32) ([]string, error) {
	depGraph := graphll.New()
	for sname, sdata := range servs.ServicesData {
		depGraph.Add(sname, sdata.Weight, sdata.Dependencies)
	}
	return graphdeppicker.Run(depGraph, quantity)
}

// PickRandDepsYml picks random nodes described in a yaml file
func PickRandDepsYml(filename string, quantity uint32) ([]string, error) {
	servs, err := getConfFile(filename)
	if err != nil {
		return nil, err
	}
	return pickRandDeps(servs, quantity)
}

// PickRandDeps picks random nodes described in a yaml file as a slice of byte
func PickRandDeps(data []byte, quantity uint32) ([]string, error) {
	servs, err := getConf(data)
	if err != nil {
		return nil, err
	}
	return pickRandDeps(servs, quantity)
}
