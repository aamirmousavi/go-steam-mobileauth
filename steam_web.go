package mobileauth

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"strings"
)

// Perform a mobile login request
// Method must be GET or POST
// Returns response body
func MobileLoginRequest(queryUrl, method string, params *url.Values, cookies *cookiejar.Jar, headers *map[string]string) ([]byte, error) {
	referer := UrlCommunityBase + "/mobilelogin?oauth_client_id=DE45CD61&oauth_scope=read_profile%20write_profile%20read_client%20write_client"
	return WebRequest(queryUrl, method, params, cookies, headers, &referer)
}

func WebRequest(queryUrl, method string, params *url.Values, cookies *cookiejar.Jar, headers *map[string]string, referer *string) ([]byte, error) {
	if referer == nil {
		aux := UrlCommunityBase
		referer = &aux
	}

	client := new(http.Client)

	// Create request
	var req *http.Request
	var err error
	switch method {
	case "GET":
		if params != nil {
			if strings.Contains(queryUrl, "?") {
				queryUrl = queryUrl + "&"
			} else {
				queryUrl = queryUrl + "?"
			}
			queryUrl = queryUrl + params.Encode()
		}
		//	fmt.Printf("m: %s\tq: %s\n", method, queryUrl)
		req, err = http.NewRequest(method, queryUrl, nil)
		if err != nil {
			panic("failed to create http request")
		}
	case "POST":
		if params == nil {
			params = &url.Values{}
		}
		req, err = http.NewRequest(method, queryUrl, bytes.NewBufferString(params.Encode()))
		if err != nil {
			panic("failed to create http request")
		}
		req.Header.Set("Content-Type", "application/x-www-form-urlencoded; charset=UTF-8")
	default:
		panic("Only POST and GET requests supported")
	}

	// Set request header params
	req.Header.Set("Accept", "text/javascript, text/html, application/xml, text/xml, */*")
	req.Header.Set("User-Agent", "Mozilla/5.0 (Linux; U; Android 4.1.1; en-us; Google Nexus 4 - 4.1.1 - API 16 - 768x1280 Build/JRO03S) AppleWebKit/534.30 (KHTML, like Gecko) Version/4.0 Mobile Safari/534.30")
	req.Header.Set("Referer", *referer)

	if headers != nil {
		for key, val := range *headers {
			req.Header.Add(key, val)
		}
	}

	// Set cookies
	if cookies != nil {
		client.Jar = cookies
	}

	// Make request
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		s, e := ioutil.ReadAll(resp.Body)
		if e != nil {
			fmt.Println("am I joke to you?")
		} else {
			Savet(queryUrl + " " + fmt.Sprintf("%v", *params) + " " + string(s))
		}
		return nil, fmt.Errorf("request failed with status code: %v", resp.StatusCode)
	}
	return ioutil.ReadAll(resp.Body)
}

//cookies
func DeleteMe(cookies []*http.Cookie) {
	for _, cook := range cookies {
		fmt.Println("cook: ", *cook)
	}
}

func WebRequestRaw(queryUrl, method string, params *url.Values, cookies *cookiejar.Jar, headers *map[string]string, referer *string) (*http.Response, error) {
	if referer == nil {
		aux := UrlCommunityBase
		referer = &aux
	}

	client := new(http.Client)

	// Create request
	var req *http.Request
	var err error
	switch method {
	case "GET":
		if params != nil {
			if strings.Contains(queryUrl, "?") {
				queryUrl = queryUrl + "&"
			} else {
				queryUrl = queryUrl + "?"
			}
			queryUrl = queryUrl + params.Encode()
		}
		//	fmt.Printf("m: %s\tq: %s\n", method, queryUrl)
		req, err = http.NewRequest(method, queryUrl, nil)
		if err != nil {
			panic("failed to create http request")
		}
	case "POST":
		if params == nil {
			params = &url.Values{}
		}
		req, err = http.NewRequest(method, queryUrl, bytes.NewBufferString(params.Encode()))
		if err != nil {
			panic("failed to create http request")
		}
		req.Header.Set("Content-Type", "application/x-www-form-urlencoded; charset=UTF-8")
	default:
		panic("Only POST and GET requests supported")
	}

	// Set request header params
	req.Header.Set("Accept", "text/javascript, text/html, application/xml, text/xml, */*")
	req.Header.Set("User-Agent", "Mozilla/5.0 (Linux; U; Android 4.1.1; en-us; Google Nexus 4 - 4.1.1 - API 16 - 768x1280 Build/JRO03S) AppleWebKit/534.30 (KHTML, like Gecko) Version/4.0 Mobile Safari/534.30")
	req.Header.Set("Referer", *referer)

	if headers != nil {
		for key, val := range *headers {
			req.Header.Add(key, val)
		}
	}

	// Set cookies
	if cookies != nil {
		client.Jar = cookies
	}

	// Make request
	return client.Do(req)

}
