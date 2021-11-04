package main

import (
	"bytes"
	b64 "encoding/base64"
	"encoding/json"

	// "errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"

	// "reflect"
	"strconv"
	"strings"
)

func Auth() {
	// start to bulid request
	form := url.Values{}
	form.Set("grant_type", "refresh_token")
	form.Set("refresh_token", RefreshToken)

	client := &http.Client{}
	request, err := http.NewRequest("POST", "https://accounts.spotify.com/api/token", strings.NewReader(form.Encode()))
	if err != nil {
		log.Fatal(err)
	}

	request.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	request.Header.Add("Authorization", "Basic "+b64.StdEncoding.EncodeToString([]byte(BasicCreds)))

	// make request
	res, err := client.Do(request)
	if err != nil {
		log.Fatal(err)
	}
	if res.StatusCode != 200 && res.StatusCode != 201 {
		log.Println("**** Failed to aquire new auth **** response code: " + strconv.Itoa(res.StatusCode))
		log.Fatal(res)
	}

	defer res.Body.Close()

	// convert body to JSON
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	body_map := make(map[string]interface{})

	err = json.Unmarshal([]byte(body), &body_map)
	if err != nil {
		panic(err)
	}

	// I don't know why we have to cast it this way, but the other way to cast didn't work
	AccessToken = body_map["access_token"].(string)
	fmt.Println("Access: "+AccessToken)
}


func BuildRequest(request_type string, url string, body []byte) (map[string]interface{}, error) {
	client := &http.Client{}

	var request *http.Request
	var err error

	if body == nil {
		request, err = http.NewRequest(request_type, url, nil)
	} else {
		request, err = http.NewRequest(request_type, url, bytes.NewBuffer(body))
	}

	if err != nil {
		log.Fatal(err)
	}
	request.Header.Add("Authorization", "Bearer "+AccessToken)
	request.Header.Set("Content-Type", "application/json; charset=UTF-8")

	fmt.Println("Sending " + request_type + "request")
	res, err := client.Do(request)
	if err != nil {
		log.Fatal(err)
	}

	// if token is expired, authenticate then run the same request again
	if res.StatusCode > 399 {
		log.Println("Auth credentials expired, requesting new token...")
		// log.Println("UNSUCCESSFUL " + strconv.Itoa(res.StatusCode) + ": " + request_type + " request for " + url)
		Auth()

		// I really wanted this to be recursive, like this
		// return BuildRequest(request_type, url, body)
		// but it's not working for some reason and I don't feel like dealing with it

		if body == nil {
			request, err = http.NewRequest(request_type, url, nil)
		} else {
			request, err = http.NewRequest(request_type, url, bytes.NewBuffer(body))
		}
	
		if err != nil {
			log.Fatal(err)
		}
		request.Header.Add("Authorization", "Bearer "+AccessToken)
		request.Header.Set("Content-Type", "application/json; charset=UTF-8")
	
		fmt.Println("Sending " + request_type + "request")
		res, err = client.Do(request)
		if err != nil {
			log.Fatal(err)
		}

	} else if res.StatusCode != 200 && res.StatusCode != 201 {
		// Literally no idea why this would happen
		log.Fatal("UNSUCCESSFUL in BuildRequest" + strconv.Itoa(res.StatusCode) + ": " + request_type + " request for " + url)
	}

	defer res.Body.Close()

	res_body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Println(err)
	}

	// convert to usable map
	body_map := make(map[string]interface{})
	err = json.Unmarshal([]byte(res_body), &body_map)
	if err != nil {
		log.Println(err)
	}

	fmt.Println(request_type + " request finished")
	return body_map, nil
}