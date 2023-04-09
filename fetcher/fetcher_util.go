package fetcher

import (
	"net/http"
)

var (
	defaultConfig = Config{
		Timeout: DefaultTimeOutSeconds,
	}
	defaultFetcher = NewFetcher(defaultConfig)
)

func Get(url string) (*http.Response, error) {
	return defaultFetcher.Get(url)
}

func GetJson[R any](url string, r R) (R, *http.Response, error) {
	res, err := defaultFetcher.GetJson(url, r)
	return r, res, err
}

func FetchGet(f *Fetcher, url string) (*http.Response, error) {
	return f.Get(url)
}

func FetchGetJson[R any](f *Fetcher, url string, r R) (R, *http.Response, error) {
	res, err := f.GetJson(url, r)
	return r, res, err
}

func Delete(url string) (*http.Response, error) {
	return defaultFetcher.Get(url)
}

func DeleteJson[R any](url string, r R) (R, *http.Response, error) {
	res, err := defaultFetcher.GetJson(url, r)
	return r, res, err
}

func FetchDelete(f *Fetcher, url string) (*http.Response, error) {
	return f.Get(url)
}

func FetchDeleteJson[R any](f *Fetcher, url string, r R) (R, *http.Response, error) {
	res, err := f.GetJson(url, r)
	return r, res, err
}

func Post[B any](url string, body B) (*http.Response, error) {
	return defaultFetcher.Post(url, body)
}

func PostJson[B any, R any](url string, body B, r R) (R, *http.Response, error) {
	res, err := defaultFetcher.PostJson(url, body, r)
	return r, res, err
}

func FetchPost[B any](f *Fetcher, url string, body B) (*http.Response, error) {
	return f.Post(url, body)
}

func FetchPostJson[B any, R any](f *Fetcher, url string, body B, r R) (R, *http.Response, error) {
	res, err := f.PostJson(url, body, r)
	return r, res, err
}

func Put[B any](url string, body B) (*http.Response, error) {
	return defaultFetcher.Post(url, body)
}

func PutJson[B any, R any](url string, body B, r R) (R, *http.Response, error) {
	res, err := defaultFetcher.PostJson(url, body, r)
	return r, res, err
}

func FetchPut[B any](f *Fetcher, url string, body B) (*http.Response, error) {
	return f.Post(url, body)
}

func FetchPutJson[B any, R any](f *Fetcher, url string, body B, r R) (R, *http.Response, error) {
	res, err := f.PostJson(url, body, r)
	return r, res, err
}
