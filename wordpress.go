package wordpress

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

const (
	apiName    string = "WordPress"
	apiPath    string = "wp-json/wp"
	apiVersion string = "v2"
)

// WordPress stores WordPress configuration
//
type WordPress struct {
	domain         string
	basicAuthToken string
	isLive         bool
}

type NewWordPressParams struct {
	Domain         string
	BasicAuthToken string
	IsLive         bool
}

// NewWordPress return new instance of WordPress struct
//
func NewWordPress(params NewWordPressParams) (*WordPress, error) {
	return &WordPress{params.Domain, params.BasicAuthToken, params.IsLive}, nil
}

func (wp *WordPress) BaseURL() string {
	return fmt.Sprintf("%s/%s/%s", wp.domain, apiPath, apiVersion)
}

// Get //
//
func (wp *WordPress) Get(url string, model interface{}) error {
	client := &http.Client{}

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return err
	}
	req.Header.Set("accept", "application/json")
	req.Header.Set("authorization", fmt.Sprintf("Basic %s", wp.basicAuthToken))

	// Send out the HTTP request
	res, err := client.Do(req)
	if err != nil {
		return err
	}

	defer res.Body.Close()

	b, err := ioutil.ReadAll(res.Body)

	errr := json.Unmarshal(b, &model)
	if errr != nil {
		return err
	}

	return nil
}
