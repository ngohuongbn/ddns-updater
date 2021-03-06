package main

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"strconv"
	"time"

	"github.com/kyokomi/emoji"
)

func connectivityTest() {
	_, err := net.LookupIP("google.com")
	if err != nil {
		log.Println(emoji.Sprint(":signal_strength:") + "Domain name resolution " + emoji.Sprint(":x:") + " is not working for google.com (" + err.Error() + ")")
	} else {
		log.Println(emoji.Sprint(":signal_strength:") + "Domain name resolution " + emoji.Sprint(":heavy_check_mark:"))
	}
	req, err := buildHTTPGet("https://google.com")
	if err != nil {
		log.Println(emoji.Sprint(":signal_strength:") + "HTTP GET " + emoji.Sprint(":x:") + " " + err.Error())
	}
	status, _, err := doHTTPRequest(req, httpGetTimeout)
	if err != nil {
		log.Println(emoji.Sprint(":signal_strength:") + "HTTP GET " + emoji.Sprint(":x:") + " " + err.Error())
	} else if status != "200" {
		log.Println(emoji.Sprint(":signal_strength:") + "HTTP GET " + emoji.Sprint(":x:") + " " + " HTTP status " + status)
	} else {
		log.Println(emoji.Sprint(":signal_strength:") + "HTTP GET " + emoji.Sprint(":heavy_check_mark:"))
	}
}

func buildHTTPGet(url string) (request *http.Request, err error) {
	request, err = http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}
	return request, nil
}

// GoDaddy
func buildHTTPPutJSONAuth(url, authorizationHeader string, body interface{}) (request *http.Request, err error) {
	jsonData, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}
	request, err = http.NewRequest(http.MethodPut, url, bytes.NewBuffer(jsonData))
	if err != nil {
		return nil, err
	}
	request.Header.Set("Authorization", authorizationHeader)
	request.Header.Set("Content-Type", "application/json; charset=utf-8")
	return request, nil
}

func doHTTPRequest(request *http.Request, timeout int) (status string, content []byte, err error) {
	client := &http.Client{Timeout: time.Duration(timeout) * time.Millisecond}
	response, err := client.Do(request)
	if err != nil {
		return status, nil, err
	}
	status = strconv.FormatInt(int64(response.StatusCode), 10)
	content, err = ioutil.ReadAll(response.Body)
	response.Body.Close()
	if err != nil {
		return status, nil, err
	}
	return status, content, nil
}
