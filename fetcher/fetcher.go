package fetcher

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"reflect"
	"strings"
	"time"
)

type Fetcher struct {
	BaseConfig Config
	BaseClient *http.Client
}

const (
	DefaultTimeOutSeconds = 30
)

func NewFetcher(c Config) *Fetcher {
	config := c.GetCopy()
	return &Fetcher{BaseConfig: config, BaseClient: newClient(config)}
}

func (c *Fetcher) GetJson(url string, result any) (*http.Response, error) {
	// 1.Do request
	res, err := c.Get(url)
	if err != nil {
		return res, err
	}
	// 2.Read Data from res
	return res, readJsonFromResponse(res, result)
}

func (c *Fetcher) Get(url string) (*http.Response, error) {
	// 1.Create a new config
	config := c.BaseConfig.GetCopy()
	config.Url = url
	config.Method = http.MethodGet
	// 2.Do request
	return c.Do(config)
}

func (c *Fetcher) PostJson(url string, body interface{}, result any) (*http.Response, error) {
	// 1.Do request
	res, err := c.Post(url, body)
	if err != nil {
		return res, err
	}
	// 2.Read Data from res
	return res, readJsonFromResponse(res, result)
}

func (c *Fetcher) Post(url string, body interface{}) (*http.Response, error) {
	// 1.Create a new config
	config := c.BaseConfig.GetCopy()
	config.Url = url
	config.Method = http.MethodPost
	config.AddHeader("Content-Type", "application/json")
	// 2.Create body
	bodyReader, err := getJsonBodyReader(body)
	if err != nil {
		return nil, err
	}
	config.Body = bodyReader
	// 3.Do request
	return c.Do(config)
}

func (c *Fetcher) PutJson(url string, body interface{}, result any) (*http.Response, error) {
	// 1.Do request
	res, err := c.Post(url, body)
	if err != nil {
		return res, err
	}
	// 2.Read Data from res
	return res, readJsonFromResponse(res, result)
}

func (c *Fetcher) Put(url string, body interface{}) (*http.Response, error) {
	// 1.Create a new config
	config := c.BaseConfig.GetCopy()
	config.Url = url
	config.Method = http.MethodPut
	config.AddHeader("Content-Type", "application/json")
	// 2.Create body
	bodyReader, err := getJsonBodyReader(body)
	if err != nil {
		return nil, err
	}
	config.Body = bodyReader
	// 3.Do request
	return c.Do(config)
}

func (c *Fetcher) DeleteJson(url string, result any) (*http.Response, error) {
	// 1.Do request
	res, err := c.Get(url)
	if err != nil {
		return res, err
	}
	// 2.Read Data from res
	return res, readJsonFromResponse(res, result)
}

func (c *Fetcher) Delete(url string) (*http.Response, error) {
	// 1.Create a new config
	config := c.BaseConfig.GetCopy()
	config.Url = url
	config.Method = http.MethodDelete
	// 2.Do request
	return c.Do(config)
}

func (c *Fetcher) Do(cfg Config) (*http.Response, error) {
	// 1.Create Request
	requestUrl := getRequestUrl(cfg.BaseUrl, cfg.Url, cfg.Params)
	log.Printf("Send Request Method: %s, Url: %s \n", cfg.Method, requestUrl)
	req, err := http.NewRequest(cfg.Method, requestUrl, cfg.Body)
	if err != nil {
		log.Printf("Create Request Failure: %v \n", err)
		return nil, err
	}
	// 2.Add header
	appendHeader(req, cfg.Headers)
	// 3.Add auth
	appendBaseAuth(req, cfg.Auth)
	// 4.Do Request
	res, err := c.BaseClient.Do(req)
	if err != nil {
		log.Printf("Send request failure: %v", err)
		return nil, err
	}
	return res, nil
}

func getRequestUrl(baseUrl string, urlPath string, params map[string]string) string {
	u := ""
	// Append base url
	if len(baseUrl) > 0 {
		u += baseUrl
	}
	// Append url path
	if len(urlPath) > 0 {
		u += urlPath
	}
	// Append params
	if params != nil && len(params) > 0 {
		if !strings.Contains(u, "?") {
			u += "?"
		}
		for k, v := range params {
			u += k + "=" + url.QueryEscape(v)
		}
	}
	return u
}

func newClient(c Config) *http.Client {
	return &http.Client{
		Timeout: getTimeout(c) * time.Second,
	}
}

func getTimeout(c Config) time.Duration {
	timeout := c.Timeout
	if timeout <= 0 {
		timeout = DefaultTimeOutSeconds
	}
	return time.Duration(timeout)
}

func appendHeader(r *http.Request, headers map[string]string) {
	if headers != nil {
		for k, v := range headers {
			r.Header.Set(k, v)
		}
	}
}

func appendBaseAuth(req *http.Request, auth Auth) {
	if reflect.ValueOf(auth).IsZero() {
		req.SetBasicAuth(auth.Username, auth.Password)
	}
}

func getJsonBodyReader(data interface{}) (io.Reader, error) {
	if data != nil {
		byteData, err := json.Marshal(data)
		if err != nil {
			log.Printf("Serialize json data with struct data failure: %v \n", err)
			return nil, err
		}
		return bytes.NewReader(byteData), err
	}
	return nil, nil
}

func readJsonFromResponse(res *http.Response, result any) error {
	// 1.Handle close res
	defer func(Body io.ReadCloser) {
		if err := Body.Close(); err != nil {
			log.Printf("Close response body failure: %v \n", err)
		}
	}(res.Body)
	// 2.Read json from res
	if res.StatusCode >= 200 && res.StatusCode < 300 {
		if err := json.NewDecoder(res.Body).Decode(result); err != nil {
			log.Printf("Deserialize json data from response failure: %v \n", err)
			return err
		}
		return nil
	} else {
		resData, err := ioutil.ReadAll(res.Body)
		if err != nil {
			log.Printf("Read all data from res failure: %v \n", err)
		}
		return errors.New(fmt.Sprintf("Response status: %s, data: %s", res.Status, string(resData)))
	}
}
