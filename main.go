package main

import (
	"encoding/json"
	"flag"
	"fmt"
	// "github.com/davecgh/go-spew/spew"
	"strconv"
)

type Statistics struct {
	ViewCount string `json:"viewCount"`
}

type Item struct {
	Statistics Statistics `json:"statistics"`
}

type VideoListResponse struct {
	Items []Item `json:"items"`
}

func main() {
	idPtr := flag.String("id", "", "id of video")
	apiKeyPtr := flag.String("key", "", "your api key")
	flag.Parse()
	//TODO: panic if no id or key are provided
	// rawJson := GetLocalJson(*idPtr, *apiKeyPtr)
	rawJson := GetYoutubeJson(*idPtr, *apiKeyPtr)

	videoListResponse := new(VideoListResponse)
	err := json.Unmarshal(rawJson, &videoListResponse)

	if err != nil {
		panic(err.Error())
	}

	viewCount, err := strconv.Atoi(videoListResponse.Items[0].Statistics.ViewCount)

	if err != nil {
		panic(err.Error())
	}

	fmt.Printf("%v\n", viewCount)
	// spew.Dump(videoListResponse)
}
