package bobune

import (
	"io"
	"log"
	"net/http"
	"os"

	"github.com/fogleman/gg"
	"github.com/nlopes/slack"
)

func CreateBobuneImg(user *slack.User) (string, error) {
	userName := user.Profile.DisplayName
	if len(userName) == 0 {
		userName = user.Profile.RealName
	}
	fileName := user.ID + ".png"
	iconPath := dlIconImage(user, fileName)

	bgImg, err := gg.LoadImage("assets/bobune.jpg")
	if err != nil {
		return "", err
	}
	width := bgImg.Bounds().Dx()
	height := bgImg.Bounds().Dy()

	iconImg, err := gg.LoadImage(iconPath)
	if err != nil {
		return "", err
	}

	dc := gg.NewContext(width, height)
	dc.SetRGB(0, 0, 0)
	dc.LoadFontFace("assets/fonts/GenShinGothic.ttf", 40)
	dc.DrawRoundedRectangle(0, 0, 512, 512, 0)
	dc.DrawImage(bgImg, 0, 0)
	dc.DrawImage(iconImg, width*4/5-96, height*1/4)
	dc.DrawStringAnchored("@"+userName, float64(width*4/5), float64(height*4/5), 0.5, 0)
	dc.Clip()
	dc.SavePNG("assets/out/" + fileName)
	return fileName, nil
}

func dlIconImage(user *slack.User, fileName string) string {
	response, err := http.Get(user.Profile.Image192)
	if err != nil {
		log.Fatal(err)
	}

	defer response.Body.Close()

	//open a file for writing
	iconPath := "assets/icons/" + fileName
	file, err := os.Create(iconPath)
	if err != nil {
		log.Fatal(err)
	}
	_, err = io.Copy(file, response.Body)
	if err != nil {
		log.Fatal(err)
	}
	file.Close()
	return iconPath
}
