package comm

import (
	"fmt"
	"net/http"
	"time"
)

var HTTPClient *http.Client

func init() {
	transport := &http.Transport{
		MaxIdleConnsPerHost: 10,
		IdleConnTimeout:     20 * time.Second,
	}
	HTTPClient = &http.Client{
		Transport: transport,
		Timeout:   10 * time.Second,
	}
}

func GetHTTPClient() *http.Client {
	return HTTPClient
}

func HttpGet(url string) (*http.Response, error) {
	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Println(err)
	}
	// f2 12 14
	request.Header.Add("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10.15; rv:81.0) Gecko/20100101 Firefox/81.0")
	return GetHTTPClient().Do(request)
}