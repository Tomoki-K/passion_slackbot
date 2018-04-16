package image

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"os"
	"strings"
	"time"
	// "github.com/Tomoki-K/passion_tarinai/models"
)

func GetImgUrl(keyword string) (string, error) {
	word := strings.TrimSpace(keyword)
	if len(word) < 1 {
		return "", errors.New("invalid search param")
	}
	baseUrl := "https://www.googleapis.com/customsearch/v1"
	s := Search{
		Key:      os.Getenv("GOOGLE_PASSION_KEY"),
		EngineId: os.Getenv("CSE_ID"),
		Type:     "image",
		Count:    "5",
	}

	url := baseUrl + "?key=" + s.Key + "&cx=" + s.EngineId + "&searchType=" + s.Type + "&num=" + s.Count + "&q=" + word
	log.Println(url)

	imageUrl := ParseJson(url)
	return imageUrl, nil
}

func ParseJson(url string) string {
	var imageUrl = "image not found"

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
		imageUrl = data.Items[rand.Intn(5)].Link
		log.Println(imageUrl)
	}

	return imageUrl
}
