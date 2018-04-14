package main

import (
	"encoding/base64"
	"log"
	"math/rand"
	"os"
	"strings"

	"github.com/nlopes/slack"
)

var aa = "" +
	"44CA44CA44CAIF8g44CA4oipDQrjgIDjgIAo44C" +
	"A776f4oiA776fKeW9oeOAgOOBiuOBo+OBseOBhC" +
	"HjgYrjgaPjgbHjgYQhDQrjgIDjgIAo44CAIOKKg" +
	"uW9oQ0K44CAIOOAgHzjgIDjgIDjgIB8DQrjgIAg" +
	"44CA44GXIOKMku+8qg=="

func IncludesPassion(text string) (bool, error) {
	keywords := [...]string{"パッション", "ぱっしょん", "passion", "Passion"}
	for _, e := range keywords {
		if strings.Contains(text, e) {
			return true, nil
		}
	}
	return false, nil
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
				isPassion, _ := IncludesPassion(ev.Text)
				if isPassion {
					// mention sender
					text := "<@" + ev.User + "> パッションが足りません。"
					if rand.Intn(100) < 5 {
						text = "<@" + ev.User + "> 温かいし止まらない。"
					}
					rtm.SendMessage(rtm.NewOutgoingMessage(text, ev.Channel))
				} else {
					if strings.Contains(ev.Text, "おっぱい") {
						b, _ := base64.StdEncoding.DecodeString(aa)
						text := "<@" + ev.User + ">\n" + string(b)
						rtm.SendMessage(rtm.NewOutgoingMessage(text, ev.Channel))
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
