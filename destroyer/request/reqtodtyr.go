package request

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

type nodesResJson struct {
	Nodes []string `json:"nodes"`
}

var ErrUnauthorized = errors.New("Invalid credentials")

// ReqToDestroy requests the chaos server to shut down nodes
func ReqToDestroy(host string, nodes []string, token string) error {
	reqBody, err := json.Marshal(nodesResJson{nodes})
	if err != nil {
		return err
	}

	url := fmt.Sprintf("http://%v/shutdown", host)
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(reqBody))
	if err != nil {
		return err
	}
	req.AddCookie(&http.Cookie{Name: "token", Value: token})
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode == 500 {
		return errors.New("Fail to destroy nodes")
	} else if resp.StatusCode == 401 {
		return ErrUnauthorized
	} else if resp.StatusCode > 401 {
		return errors.New("Fail to communicate with chaos server")
	}

	return nil
}
