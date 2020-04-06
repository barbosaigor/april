package request

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/barbosaigor/april"
	"github.com/barbosaigor/april/destroyer"
)

var ErrUnauthorized = errors.New("Invalid credentials")

// ReqToDestroy requests the chaos server to shut down services
func ReqToDestroy(host string, svcs []april.Service, token string) error {
	// Create http body request
	svcsBody := make([]destroyer.ServiceBodyJson, len(svcs))
	for i, svc := range svcs {
		svcsBody[i].Name = svc.Name
		svcsBody[i].Selector = svc.Selector
	}
	reqBody, err := json.Marshal(destroyer.ShutdownBodyJson{svcsBody})
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
		return errors.New("Fail to destroy services")
	} else if resp.StatusCode == 401 {
		return ErrUnauthorized
	} else if resp.StatusCode > 401 {
		return errors.New("Fail to communicate with chaos server")
	}

	return nil
}
