package destroyer

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/barbosaigor/april"
	"github.com/barbosaigor/april/auth"
	"github.com/barbosaigor/april/selector"
)

var ErrNotMatchingService = errors.New("Some service has no instance matching")

type Destroyer struct {
	ChaosSrv ChaosServer
}

// getInstancesFromSvcs selects all matched services from instances (using every service notation)
func (d *Destroyer) getInstancesFromSvcs(instances []Instance, svcs []april.Service) []string {
	// Use map force unique key, even either multiple services matches
	itcs := make(map[string]string, len(svcs))
	for _, instance := range instances {
		for _, svc := range svcs {
			// Set a default selector
			sel, ok := selector.Selector[svc.Selector]
			if !ok {
				sel = selector.All
			}
			if selector.Match(instance.Name, svc.Name, sel) {
				itcs[svc.Name] = instance.Name
			}
		}
	}
	var instcs []string
	for _, name := range itcs {
		instcs = append(instcs, name)
	}
	return instcs
}

// Shutdown turns down services listed on svcs
func (d *Destroyer) Shutdown(svcs []april.Service) error {
	instances, err := d.ChaosSrv.ListInstances(Up)
	if err != nil {
		return err
	}
	itcs := d.getInstancesFromSvcs(instances, svcs)
	if len(svcs) > len(itcs) {
		return ErrNotMatchingService
	}
	for _, instance := range itcs {
		if err = d.ChaosSrv.Shutdown(instance); err != nil {
			return err
		}
	}
	return nil
}

type server struct {
	Cred     *auth.Credentials
	port     int
	destyer  Destroyer
	serveMux *http.ServeMux
}

type ServiceBodyJson struct {
	Name     string `json:"name"`
	Selector string `json:"selector"`
}

type ShutdownBodyJson struct {
	Services []ServiceBodyJson `json:"services"`
}

type ResponseMessage struct {
	Message string `json:"message"`
}

func New(port int, cs ChaosServer) *server {
	return &server{auth.New(), port, Destroyer{cs}, nil}
}

// shutDownHandler shut down instances
// body:
//		nodes: nodes to shut down
func (s *server) shutDownHandler() http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		switch r.Method {
		case "POST":
			data, err := ioutil.ReadAll(r.Body)
			if err != nil {
				resMsg := ResponseMessage{"Error to read request"}
				res, _ := json.Marshal(resMsg)
				w.WriteHeader(http.StatusInternalServerError)
				w.Write(res)
				return
			}

			var sbjson ShutdownBodyJson
			err = json.Unmarshal(data, &sbjson)
			if err != nil {
				resMsg := ResponseMessage{"Invalid request body. Should be a valid json"}
				res, _ := json.Marshal(resMsg)
				w.WriteHeader(http.StatusInternalServerError)
				w.Write(res)
				return
			}

			// Create services using services data from request
			svcs := make([]april.Service, len(sbjson.Services))
			for i, svc := range sbjson.Services {
				svcs[i] = april.Service{svc.Name, svc.Selector}
			}

			err = s.destyer.Shutdown(svcs)
			if err != nil {
				resMsg := ResponseMessage{err.Error()}
				res, _ := json.Marshal(resMsg)
				w.WriteHeader(http.StatusInternalServerError)
				w.Write(res)
				return
			}
		default:
			w.WriteHeader(http.StatusForbidden)
			w.Write([]byte("{}"))
		}
	})
}

// Serve hosts aprils API over HTTP protocol
func (s *server) Serve() {
	s.serveMux = http.NewServeMux()
	s.serveMux.Handle("/shutdown", s.Cred.MwAuth(s.shutDownHandler()))
	fmt.Println("(HTTP) Listening on port: ", s.port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%v", s.port), s.serveMux))
}
