package destroyer

import (
	"fmt"
	"log"
	"net/http"
	"io/ioutil"
	"encoding/json"
)

type Destroyer interface {
	// Destroy does shutdown nodes
	Destroy(nodes []string) error
}

type Server struct {
	Port int
	Destyer Destroyer
	
	serveMux *http.ServeMux
}

type nodesJson struct {
	Nodes []string `json:"nodes"`
}

type responseMessage struct {
	Message string `json:"message"`
}

// mwSetJsonHeader sets the http header to support json responses
func mwSetJsonHeader(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		next.ServeHTTP(w, r)
	})
}

// shutDownHandler shut down instances
// body:
//		nodes: nodes to shut down
func (s *Server) shutDownHandler() http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
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

			var njson nodesJson
			err = json.Unmarshal(data, &njson)
			if err != nil {
				resMsg := responseMessage{"Invalid request body. Should be a valid json"}
				res, _ := json.Marshal(resMsg)
				w.WriteHeader(http.StatusInternalServerError)
				w.Write(res)
				return
			}

			err = s.Destyer.Destroy(njson.Nodes)
			if err != nil {
				resMsg := responseMessage{"One container had a problem"}
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
	s.serveMux = http.NewServeMux()
	s.serveMux.Handle("/stop", mwSetJsonHeader(s.shutDownHandler()))
	fmt.Println("(HTTP) Listening on port: ", s.Port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%v", s.Port), s.serveMux))
}
