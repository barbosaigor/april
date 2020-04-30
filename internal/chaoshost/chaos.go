package chaoshost

import (
	"github.com/barbosaigor/april"
	"github.com/barbosaigor/april/destroyer/request"
	"github.com/barbosaigor/april/util"
)

// ChaosHost stores some server data and run chaos test operations
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

// PickAndShutdownInstances selects a set of services and request to chaos server to shut down
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

// PickAndShutdownInstancesFile as PickAndShutdownInstances does, but the it reads from a file the services data
func (ch ChaosHost) PickAndShutdownInstancesFile(filePath string, n uint32) ([]string, error) {
	conf, err := getConf(filePath)
	if err != nil {
		return nil, err
	}
	return ch.PickAndShutdownInstances(conf, n)
}
