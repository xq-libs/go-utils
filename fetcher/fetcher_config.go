package fetcher

import "io"

type Config struct {
	Url     string
	BaseUrl string
	Method  string
	Headers map[string]string
	Params  map[string]string
	Body    io.Reader
	Auth    Auth
	Timeout int64
}

type Auth struct {
	Username string
	Password string
}

func (c *Config) AddHeader(key string, value string) {
	if c.Headers == nil {
		c.Headers = make(map[string]string, 0)
	}
	c.Headers[key] = value
}

func (c *Config) AddParam(key string, value string) {
	if c.Params == nil {
		c.Params = make(map[string]string, 0)
	}
	c.Params[key] = value
}

func (c *Config) GetCopy() Config {
	return Config{
		Url:     c.Url,
		BaseUrl: c.BaseUrl,
		Method:  c.Method,
		Headers: copyMap(c.Headers),
		Params:  copyMap(c.Params),
		Body:    c.Body,
		Auth:    c.Auth,
		Timeout: c.Timeout,
	}
}

func copyMap(m map[string]string) map[string]string {
	if m == nil {
		return m
	}
	rm := make(map[string]string, len(m))
	for k, v := range m {
		rm[k] = v
	}
	return rm
}
