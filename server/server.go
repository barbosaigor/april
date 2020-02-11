package server

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"io/ioutil"
	"encoding/json"

	"gitlab.com/barbosaigor/april"
	"gitlab.com/barbosaigor/april/destroyer/request"
)

const defaultDestroyerHost = "localhost:7071"

type confResJson struct {
	Conf string `json:"conf"`
	DestroyerHost string `json:"destroyerHost"`
}

type nodesResJson struct {
	Nodes []string `json:"nodes"`
}

// bareHandler executes only the node selecting algorithm 
// query:
//		n is the number of returning nodes
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

	nodes, err := april.PickRandDeps([]byte(c.Conf), uint32(n))
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Fail to pick nodes", http.StatusInternalServerError)
		return
	}

	nRes, _ := json.Marshal(nodesResJson{nodes})
	w.Header().Set("Content-Type", "application/json")
	w.Write(nRes)
}

// chaosHandler apply chaos testing
// query:
//		n is the number of returning nodes
// body:
//		conf is the configuration file (yaml file)
//		destroyerHost is the endpoint of destroyer server
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
	if c.DestroyerHost == "" {
		c.DestroyerHost = defaultDestroyerHost
	}
	
	nodes, err := april.PickRandDeps([]byte(c.Conf), uint32(n))
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Fail to pick nodes", http.StatusInternalServerError)
		return
	}
	
	err = request.ReqToDestroy(c.DestroyerHost, nodes)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "There was a problem with destroy server", http.StatusInternalServerError)
		return	
	}

	nRes, _ := json.Marshal(nodesResJson{nodes})
	w.Header().Set("Content-Type", "application/json")
	w.Write(nRes)
}

// Serve serves aprils API over HTTP protocol
func Serve(port int) {
	http.HandleFunc("/", chaosHandler)
	http.HandleFunc("/bare", bareHandler)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%v", port), nil))
}