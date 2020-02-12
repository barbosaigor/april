package destroyer

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/barbosaigor/april/auth"
)

type Destroyer interface {
	// Destroy does shutdown nodes
	Destroy(nodes []string) error
}

type server struct {
	Cred     *auth.Credentials
	port     int
	destyer  Destroyer
	serveMux *http.ServeMux
}

type nodesJson struct {
	Nodes []string `json:"nodes"`
}

type responseMessage struct {
	Message string `json:"message"`
}

func New(port int, destyer Destroyer) *server {
	return &server{auth.New(), port, destyer, nil}
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

			var njson nodesJson
			err = json.Unmarshal(data, &njson)
			if err != nil {
				resMsg := responseMessage{"Invalid request body. Should be a valid json"}
				res, _ := json.Marshal(resMsg)
				w.WriteHeader(http.StatusInternalServerError)
				w.Write(res)
				return
			}

			err = s.destyer.Destroy(njson.Nodes)
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
func (s *server) Serve() {
	s.serveMux = http.NewServeMux()
	s.serveMux.Handle("/shutdown", s.Cred.MwAuth(s.shutDownHandler()))
	fmt.Println("(HTTP) Listening on port: ", s.port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%v", s.port), s.serveMux))
}
