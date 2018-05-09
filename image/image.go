package image

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"net/url"
	"os"
	"time"
)

func ImageSearch(keyword string) (string, error) {
	word := &url.URL{Path: keyword}
	baseUrl := "https://www.googleapis.com/customsearch/v1"
	s := Search{
		Key:      os.Getenv("GOOGLE_PASSION_KEY"),
		EngineId: os.Getenv("CSE_ID"),
		Type:     "image",
		Count:    "5",
	}
	searchUrl := baseUrl + "?key=" + s.Key + "&cx=" + s.EngineId + "&searchType=" + s.Type + "&num=" + s.Count + "&q=" + word.String()
	log.Println("image search: " + searchUrl)
	return GetImgUrl(searchUrl), nil
}

func GetImgUrl(url string) string {
	var imgUrl = "image not found"

	response, err := http.Get(url)
	if err != nil {
		log.Println("error:", err)
	}

	defer response.Body.Close()

	byteArray, _ := ioutil.ReadAll(response.Body)

	jsonBytes := ([]byte)(byteArray)
	data := new(Result)
	if err := json.Unmarshal(jsonBytes, data); err != nil {
		log.Println("json error:", err)
	}
	if data.Items != nil {
		rand.Seed(time.Now().UTC().UnixNano())
		imgUrl = data.Items[rand.Intn(5)].Link
		log.Println(imgUrl)
	}

	return imgUrl
}
