package server

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"

	"github.com/barbosaigor/april"
	"github.com/barbosaigor/april/auth"
	"github.com/barbosaigor/april/destroyer/request"
)

var destroyerHost = "localhost:7071"

type confResJson struct {
	Conf     string `json:"conf"`
	Username string `json:"username"`
	Password string `json:"password"`
}

type resJson struct {
	Services []string `json:"services"`
}

func SetDestroyerHost(h string) {
	destroyerHost = h
}

// bareHandler executes only the service selecting algorithm
// query:
//		n is the number of returning services
// body:
//		conf is the configuration file (yaml file)
func bareHandler(w http.ResponseWriter, r *http.Request) {
	n, err := strconv.ParseUint(r.FormValue("n"), 10, 32)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	data, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Fail to read request body", http.StatusInternalServerError)
		return
	}

	var c confResJson
	json.Unmarshal(data, &c)
	if c.Conf == "" {
		http.Error(w, "Empty conf file", http.StatusInternalServerError)
		return
	}

	svs, err := april.PickRandDeps([]byte(c.Conf), uint32(n))
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Fail to pick services", http.StatusInternalServerError)
		return
	}

	nRes, _ := json.Marshal(resJson{svs})
	w.Header().Set("Content-Type", "application/json")
	w.Write(nRes)
}

// chaosHandler apply chaos testing
// query:
//		n is the number of returning services
// body:
//		conf is the configuration file (yaml file)
//		username for auth
//		password for auth
func chaosHandler(w http.ResponseWriter, r *http.Request) {
	n, err := strconv.ParseUint(r.FormValue("n"), 10, 32)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	data, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Fail to read request body", http.StatusInternalServerError)
		return
	}

	var c confResJson
	json.Unmarshal(data, &c)
	if c.Conf == "" {
		http.Error(w, "Empty conf or destroyerHost", http.StatusInternalServerError)
		return
	}

	conf, err := april.ReadConf([]byte(c.Conf))
	if err != nil {
		return
	}
	services, err := april.PickRandDepsConf(conf, uint32(n))
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Fail to pick services", http.StatusInternalServerError)
		return
	}

	token := auth.EncryptUser(c.Username, c.Password)
	svs := make([]april.Service, len(services))
	for i, s := range services {
		svs[i].Name = s
		svs[i].Selector = conf.Services[s].Selector;
	}
	err = request.ReqToDestroy(destroyerHost, svs, token)
	if err == request.ErrUnauthorized {
		http.Error(w, "Invalid user", http.StatusForbidden)
		return
	} else if err != nil {
		log.Println(err.Error())
		http.Error(w, "There was a problem with chaos server", http.StatusInternalServerError)
		return
	}
	nRes, _ := json.Marshal(resJson{services})
	w.Header().Set("Content-Type", "application/json")
	w.Write(nRes)
}

// Serve serves aprils API over HTTP protocol
func Serve(port int) {
	http.HandleFunc("/", chaosHandler)
	http.HandleFunc("/bare", bareHandler)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%v", port), nil))
}
