package tgbot

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"net/url"
	"reflect"
	"strings"
)

// Params params for telegram Bot API methods
type Params map[string]interface {
}

func reflectData(v interface{}) reflect.Value {
	rv := reflect.ValueOf(v)
	for rv.Kind() == reflect.Ptr || rv.Kind() == reflect.Interface {
		rv = rv.Elem()
	}
	return rv
}

// URLValues convert Params to url.Values, non-string values are JSON-marshaled
func (params *Params) URLValues() (url.Values, error) {
	values := url.Values{}
	for key, value := range *params {
		reflectData := reflectData(value)
		if reflectData.Kind() == reflect.String {
			values.Add(key, reflectData.String())
		} else {
			data, err := json.Marshal(value)
			if err != nil {
				return url.Values{}, errors.New("tgbotapi.Params.URLValues: " + err.Error())
			}
			values.Add(key, string(data))
		}
	}
	return values, nil
}

// URLEncode URL-encode params, returns io.Reader
func (params *Params) URLEncode() (io.Reader, error) {
	values, err := params.URLValues()
	if err != nil {
		return nil, err
	}
	return strings.NewReader(values.Encode()), nil
}

// JSONEncode URL-encode params, returns io.Reader
func (params *Params) JSONEncode() (io.Reader, error) {
	data, err := json.Marshal(params)
	if err != nil {
		return nil, err
	}
	return bytes.NewReader(data), nil
}

// MultipartFormEncode encode params as multipart/form-data, returns io.Reader
func (params *Params) MultipartFormEncode() (io.Reader, error) {
	return nil, nil
}
