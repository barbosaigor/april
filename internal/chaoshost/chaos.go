package chaoshost

import (
	"github.com/barbosaigor/april"
	"github.com/barbosaigor/april/destroyer/request"
	"github.com/barbosaigor/april/util"
)

type ChaosHost struct {
	Host  string
	Token string
}

func getConf(filePath string) (*april.ConfData, error) {
	fData, err := util.ReadFile(filePath)
	if err != nil {
		return nil, err
	}
	return april.ReadConf(fData)
}

func (ch ChaosHost) PickAndShutdownInstances(conf *april.ConfData, n uint32) ([]string, error) {
	services, err := april.PickRandDepsConf(conf, n)
	if err != nil {
		return nil, err
	}
	svs := make([]april.Service, len(services))
	for i, svcName := range services {
		svs[i].Name = svcName
		svs[i].Selector = conf.Services[svcName].Selector
	}
	if err = request.ReqToDestroy(ch.Host, svs, ch.Token); err != nil {
		return nil, err
	}
	return services, nil
}

func (ch ChaosHost) PickAndShutdownInstancesFile(filePath string, n uint32) ([]string, error) {
	conf, err := getConf(filePath)
	if err != nil {
		return nil, err
	}
	return ch.PickAndShutdownInstances(conf, n)
}
