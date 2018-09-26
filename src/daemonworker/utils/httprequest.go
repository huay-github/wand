package utils

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

func EncodeString(content string) string {
	return base64.StdEncoding.EncodeToString([]byte(content))
}

func BasicAuthEncode(username, password string) string {
	return EncodeString(username + ":" + password)
}

func DefaultResponseHandler(req *http.Request, resp *http.Response) error {
	requestURL := req.URL.String()
	if resp.StatusCode >= http.StatusBadRequest {
		return fmt.Errorf("unexpected error occurred while requesting %s with status code: %d", requestURL, resp.StatusCode)
	}
	log.Printf("Requested: %s with response status code: %d\n", requestURL, resp.StatusCode)
	return nil
}

func RequestHandle(method string, urlStr string, callback func(req *http.Request) error, data interface{}, handler func(req *http.Request, resp *http.Response) error) error {
	var payload io.Reader
	var err error
	if data != nil {
		if content, ok := data.(string); ok {
			payload = bytes.NewBuffer([]byte(content))
		} else {
			obj, err := json.Marshal(data)
			if err != nil {
				log.Printf("Failed to marshal data: %+v\n", err)
				return err
			}
			payload = bytes.NewReader(obj)
		}
	}
	req, err := http.NewRequest(method, urlStr, payload)
	if err != nil {
		return err
	}
	log.Printf("Requested URL: %s, with method: %s.\n", urlStr, method)
	if callback != nil {
		err = callback(req)
	}
	if err != nil {
		return err
	}
	client := http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if handler != nil {
		return handler(req, resp)
	}
	return DefaultResponseHandler(req, resp)
}

func DefaultRequestHandler(method, urlStr string, header http.Header, data interface{}) error {
	return RequestHandle(method, urlStr, func(req *http.Request) error {
		if header != nil {
			req.Header = header
		}
		return nil
	}, data, DefaultResponseHandler)
}

func SimpleGetRequestHandle(urlStr string) error {
	return RequestHandle(http.MethodGet, urlStr, nil, nil, DefaultResponseHandler)
}
func SimpleHeadRequestHandle(urlStr string, header http.Header) error {
	return DefaultRequestHandler(http.MethodHead, urlStr, header, nil)
}
func SimpleDeleteRequestHandle(urlStr string, header http.Header) error {
	return DefaultRequestHandler(http.MethodDelete, urlStr, header, nil)
}

func SimplePostRequestHandle(urlStr string, header http.Header, data interface{}) error {
	return DefaultRequestHandler(http.MethodPost, urlStr, header, data)
}
