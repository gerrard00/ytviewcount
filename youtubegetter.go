package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

func GetYoutubeJson(id string, apiKey string) []byte {
	// TODO: take as argument
	url := "https://www.googleapis.com/youtube/v3/videos?part=statistics&id=" +
		id + "&key=" + apiKey
	var netClient = &http.Client{
		Timeout: time.Second * 10,
	}

	response, err := netClient.Get(url)

	if err != nil {
		fmt.Printf("fuck %v\n", err)
		os.Exit(-1)
	}

	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)

	if err != nil {
		panic(err.Error())
	}
	return body
}
