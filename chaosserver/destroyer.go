package chaosserver

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/sirupsen/logrus"

	"github.com/barbosaigor/april"
	"github.com/barbosaigor/april/auth"
	"github.com/barbosaigor/april/internal/requestbody"
	"github.com/barbosaigor/april/selector"
)

// ErrNotMatchingService indicates that some selected service has no instance matching on Chaos Server
var ErrNotMatchingService = errors.New("Some selected service has no instance matching on Chaos Server")

// Destroyer implements chaos server operations
type Destroyer struct {
	ChaosSrv ChaosServer
}

// getInstancesFromSvcs selects all matched services from instances (using every service notation).
// if a service doesn't defines a selector, then Infinx will be used.
func (d *Destroyer) getInstancesFromSvcs(instances []Instance, svcs []april.Service) []string {
	// Use map force unique key, even either multiple services matches
	itcs := make(map[string]string, len(svcs))
	for _, instance := range instances {
		for _, svc := range svcs {
			// Set Infix as default selector
			sel, ok := selector.Selector[svc.Selector]
			if !ok {
				sel = selector.Infix
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

// Server implements chaos server API
type Server struct {
	Cred     *auth.Credentials
	port     int
	cs       Destroyer
	serveMux *http.ServeMux
}

// New creates a chaos server instance
func New(port int, cs ChaosServer) *Server {
	return &Server{auth.New(), port, Destroyer{cs}, nil}
}

// shutDownHandler shut down instances
// body:
//		nodes: nodes to shut down
func (s *Server) shutDownHandler() http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		switch r.Method {
		case "POST":
			data, err := ioutil.ReadAll(r.Body)
			if err != nil {
				resMsg := requestbody.ResponseMessage{Message: "Error to read request"}
				res, _ := json.Marshal(resMsg)
				w.WriteHeader(http.StatusInternalServerError)
				w.Write(res)
				return
			}

			var sbjson requestbody.ShutdownBodyJSON
			err = json.Unmarshal(data, &sbjson)
			if err != nil {
				resMsg := requestbody.ResponseMessage{Message: "Invalid request body. Should be a valid json"}
				res, _ := json.Marshal(resMsg)
				w.WriteHeader(http.StatusInternalServerError)
				w.Write(res)
				return
			}

			// Create services using services data from request
			svcs := make([]april.Service, len(sbjson.Services))
			for i, svc := range sbjson.Services {
				svcs[i] = april.Service{Name: svc.Name, Selector: svc.Selector}
			}

			err = s.cs.Shutdown(svcs)
			if err != nil {
				resMsg := requestbody.ResponseMessage{Message: err.Error()}
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
func (s *Server) Serve() {
	s.cs.ChaosSrv.OnStart()
	s.serveMux = http.NewServeMux()
	s.serveMux.Handle("/shutdown", s.Cred.MwAuth(s.shutDownHandler()))
	logrus.Fatal(http.ListenAndServe(fmt.Sprintf(":%v", s.port), s.serveMux))
}
