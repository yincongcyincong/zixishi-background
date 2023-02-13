package utils

import (
	"strings"
	"net/url"
	"net/http"
	"io/ioutil"
)

func NewRequest() *Request {
	return &Request{}
}

type Request struct {

}

// http get request
func (utils *Request) Get(queryUrl string, queryValues map[string]string, headerValues map[string]string) (body string, code int, err error) {

	req, err := http.NewRequest("GET", utils.QueryBuilder(queryUrl, queryValues), nil)
	if err != nil {
		return
	}
	if (headerValues != nil) && (len(headerValues) > 0) {
		for key, value := range headerValues {
			req.Header.Set(key, value)
		}
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	code = resp.StatusCode
	defer resp.Body.Close()

	bodyByte, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}

	return string(bodyByte), code, nil
}

// http post request
func (utils *Request) HttpPost(queryUrl string, queryValues map[string]string, headerValues map[string]string) (body string, code int, err error) {
	if !strings.Contains(queryUrl, "?") {
		queryUrl += "?"
	}
	queryString := ""
	for queryKey, queryValue := range queryValues {
		queryString = queryString + "&" + queryKey + "=" + url.QueryEscape(queryValue)
	}
	queryString = strings.Replace(queryString, "&", "", 1)

	req, err := http.NewRequest("POST", queryUrl, strings.NewReader(queryString))
	if err != nil {
		return
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	if (headerValues != nil) && (len(headerValues) > 0) {
		for key, value := range headerValues {
			req.Header.Set(key, value)
		}
	}
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return
	}
	code = resp.StatusCode
	defer resp.Body.Close()

	bodyByte, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}

	return string(bodyByte), code, nil
}

// build query params
func (utils *Request) QueryBuilder(queryUrl string, queryValues map[string]string) string {
	if !strings.Contains(queryUrl, "?") {
		queryUrl += "?"
	}

	queryString := ""
	for queryKey, queryValue := range queryValues {
		queryString = queryString + "&" + queryKey + "=" + url.QueryEscape(queryValue)
	}
	queryString = strings.Replace(queryString, "&", "", 1)
	queryUrl += queryString

	return queryUrl
}

// parse params(name=nick&pass=123)
func (utils *Request) ParseString(params string) map[string]string {

	paramsMap := map[string]string{}
	for _, param := range strings.Split(params, "&") {
		if ! strings.Contains(param, "=") {
			continue
		}
		paramList :=strings.Split(param,"=")
		paramsMap[paramList[0]] = paramList[1]
	}
	return paramsMap
}
