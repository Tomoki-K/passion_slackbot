package main

import (
	"encoding/base64"
	"log"
	"math/rand"
	"os"
	"strings"

	"github.com/nlopes/slack"
)

var passionMsg = "パッションが足りません。"
var rareMsg = "温かいし止まらない。"

var aa = "" +
	"44CA44CA44CAIF8g44CA4oipDQrjgIDjgIAo44C" +
	"A776f4oiA776fKeW9oeOAgOOBiuOBo+OBseOBhC" +
	"HjgYrjgaPjgbHjgYQhDQrjgIDjgIAo44CAIOKKg" +
	"uW9oQ0K44CAIOOAgHzjgIDjgIDjgIB8DQrjgIAg" +
	"44CA44GXIOKMku+8qg=="

var keyword2, _ = base64.StdEncoding.DecodeString("44GK44Gj44Gx44GE")

func IncludesPassion(text string) (bool, error) {
	keywords := [...]string{"パッション", "ぱっしょん", "passion", "Passion"}
	for _, e := range keywords {
		if strings.Contains(text, e) {
			return true, nil
		}
	}
	return false, nil
}

func decodeAA(encAA string) string {
	b, err := base64.StdEncoding.DecodeString(encAA)
	if err != nil {
		log.Print(err)
		return passionMsg // fallback to default if err occures
	}
	return "\n" + string(b)
}

func run(api *slack.Client) int {
	rtm := api.NewRTM()
	go rtm.ManageConnection()

	for {
		select {
		case msg := <-rtm.IncomingEvents:
			switch ev := msg.Data.(type) {
			case *slack.HelloEvent:
				log.Print("Hello Event")

			case *slack.MessageEvent:
				log.Printf("Message: %v\n", ev)
				mentionTag := "<@" + ev.User + "> "
				isPassion, _ := IncludesPassion(ev.Text)
				if isPassion {
					text := mentionTag + passionMsg // default message
					if rand.Intn(100) < 5 {
						text = mentionTag + rareMsg // 5% chance of rare message
					}
					rtm.SendMessage(rtm.NewOutgoingMessage(text, ev.Channel))
				} else {
					// other matches
					if strings.Contains(ev.Text, string(keyword2)) {
						rtm.SendMessage(rtm.NewOutgoingMessage(mentionTag+decodeAA(aa), ev.Channel))
					}
				}

			case *slack.InvalidAuthEvent:
				log.Print("Invalid credentials")
				return 1
			}
		}
	}
}

func main() {
	api := slack.New(os.Getenv("SLACK_PASSION_KEY"))
	os.Exit(run(api))
}
