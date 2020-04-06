package destroyer

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/barbosaigor/april"
	"github.com/barbosaigor/april/auth"
	"github.com/barbosaigor/april/selector"
)

type Destroyer struct {
	ChaosSrv ChaosServer
}

// FilterSvcs selects all matched services from instances (using every service notation)
func (d *Destroyer) FilterSvcs(instances []Instance, svcs []april.Service) (s []april.Service) {
	for _, instance := range instances {
		for _, svc := range svcs {
			if selector.Match(instance.Name, svc.Name, selector.Selector[svc.Selector]) {
				s = append(s, svc)
			}
		}
	}
	return
}

// Shutdown turns down services listed on svcs
func (d *Destroyer) Shutdown(svcs []april.Service) error {
	instances, err := d.ChaosSrv.ListInstances(Any)
	if err != nil {
		return err
	}
	svcs = d.FilterSvcs(instances, svcs)
	for _, svc := range svcs {
		err = d.ChaosSrv.Shutdown(svc.Name)
		if err != nil {
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

type responseMessage struct {
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
				resMsg := responseMessage{"Error to read request"}
				res, _ := json.Marshal(resMsg)
				w.WriteHeader(http.StatusInternalServerError)
				w.Write(res)
				return
			}

			var sbjson ShutdownBodyJson
			err = json.Unmarshal(data, &sbjson)
			if err != nil {
				resMsg := responseMessage{"Invalid request body. Should be a valid json"}
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
				resMsg := responseMessage{"Internal Chaos Server error"}
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
