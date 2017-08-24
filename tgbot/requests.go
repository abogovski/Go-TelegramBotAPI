package tgbot

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
)

// Available HTTP request types:
// 	URL query string
// 	application/x-www-form-urlencoded
// 	application/json (except for uploading files)
// 	multipart/form-data (use to upload files)

type request struct {
	method       string
	url          string
	status       string
	response     string
	finishedWith string
}

func newRequest(method string, url string) *request {
	st := &request{}
	st.method = method
	st.url = url
	return st
}

func (r request) log() {
	log.Printf("Request finished\n\t%v %v\n\tStatus: %v\n\tResponse: %v\n\tFinishedWith: %v\n",
		r.method, r.url, r.status, r.response, r.finishedWith)
}

func (r *request) process(httpResponse *http.Response, err error) (*Response, int, error) {
	defer func(r *request) { r.log() }(r)

	if err != nil {
		r.finishedWith = err.Error()
		return nil, httpResponse.StatusCode, errors.New("tgbot " + r.method + " request failed: " + err.Error())
	}
	r.status = httpResponse.Status

	// deserialize http response
	response := &Response{}
	decoder := json.NewDecoder(httpResponse.Body)
	err = decoder.Decode(response)
	if err != nil {
		buffered, _ := ioutil.ReadAll(decoder.Buffered())
		pending, _ := ioutil.ReadAll(httpResponse.Body)
		r.response = string(buffered) + string(pending)
		r.finishedWith = fmt.Sprintf("Failed to unmarshal TelegramBotAPI response: %v\n", err)
		return nil, httpResponse.StatusCode, errors.New("tgbot http " + r.method + " request: " + err.Error())
	}

	remarshaledResponse, _ := json.Marshal(response)
	r.response = string(remarshaledResponse)
	r.finishedWith = "Success"
	return response, httpResponse.StatusCode, nil
}

// Get GET
func Get(botAPIURL string, methodName string, params Params) (*Response, int, error) {
	url := botAPIURL + methodName
	if len(params) > 0 {
		urlValues, err := params.URLValues()
		if err != nil {
			return nil, 0, errors.New("tgbot.Get: " + err.Error())
		}
		url = url + "?" + urlValues.Encode()
	}

	return newRequest("GET", url).process(http.Get(url))
}

// Post POST
func Post(botAPIURL string, methodName string, contentType string, contentReader io.Reader) (*Response, int, error) {
	url := botAPIURL + methodName
	return newRequest("POST", url).process(http.Post(url, contentType, contentReader))
}

// PostURLEncoded POST application/x-www-form-urlencoded
func PostURLEncoded(botAPIURL string, methodName string, params Params) (*Response, int, error) {
	contentReader, err := params.URLEncode()
	if err != nil {
		return nil, 0, errors.New("tgbot.PostURLEncoded: " + err.Error())
	}
	return Post(botAPIURL, methodName, "application/x-www-form-urlencoded", contentReader)
}

// PostJSON POST application/json
func PostJSON(botAPIURL string, methodName string, params Params) (*Response, int, error) {
	contentReader, err := params.JSONEncode()
	if err != nil {
		return nil, 0, errors.New("tgbot.PostJSON: " + err.Error())
	}
	return Post(botAPIURL, methodName, "application/json", contentReader)
}

// PostMultipartForm POST multipart/form
func PostMultipartForm(botAPIURL string, methodName string, params Params) (*Response, int, error) {
	contentReader, err := params.MultipartFormEncode()
	if err != nil {
		return nil, 0, errors.New("tgbot.PostMultipartForm: " + err.Error())
	}
	return Post(botAPIURL, methodName, "multipart/form-data", contentReader)
}
