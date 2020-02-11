package request

import (
	"encoding/json"
	"net/http"
	"bytes"
	"fmt"
	"errors"
)

type nodesResJson struct {
	Nodes []string `json:"nodes"`
}

// ReqToDestroy requests the chaos server to shut down nodes
func ReqToDestroy(host string, nodes []string) error {
	reqBody, err := json.Marshal(nodesResJson{nodes})
	if err != nil {
		return err
	}

	// Make a request
	resp, err := http.Post(fmt.Sprintf("http://%v/stop", host),
		"application/json", bytes.NewBuffer(reqBody))
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode == 500 {
		return errors.New("Fail to destroy nodes")
	} else if resp.StatusCode >= 400 {
		return errors.New("Fail to communicate with chaos server")
	}

	return nil
}