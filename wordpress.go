package wordpress

import (
	"encoding/json"
	"fmt"
	errortools "github.com/leapforce-libraries/go_errortools"
	"io/ioutil"
	"net/http"
)

const (
	apiName    string = "Service"
	apiPath    string = "wp-json/wp"
	apiVersion string = "v2"
)

// Service stores Service configuration
//
type Service struct {
	domain         string
	basicAuthToken string
}

type ServiceConfig struct {
	Domain         string
	BasicAuthToken string
}

func NewService(cfg *ServiceConfig) (*Service, error) {
	return &Service{cfg.Domain, cfg.BasicAuthToken}, nil
}

func (wp *Service) BaseURL() string {
	return fmt.Sprintf("%s/%s/%s", wp.domain, apiPath, apiVersion)
}

func (wp *Service) Get(url string, model interface{}) *errortools.Error {
	client := &http.Client{}

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return errortools.ErrorMessage(err)
	}
	req.Header.Set("accept", "application/json")
	req.Header.Set("authorization", fmt.Sprintf("Basic %s", wp.basicAuthToken))

	// Send out the HTTP request
	res, err := client.Do(req)
	if err != nil {
		return errortools.ErrorMessage(err)
	}

	defer res.Body.Close()

	b, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return errortools.ErrorMessage(err)
	}

	err = json.Unmarshal(b, &model)
	if err != nil {
		return errortools.ErrorMessage(err)
	}

	return nil
}
