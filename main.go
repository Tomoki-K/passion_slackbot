package main

import (
	"log"
	"os"
	"strings"

	"github.com/nlopes/slack"
)

func IncludesPassion(text string) bool {
	keywords := [...]string{"パッション", "ぱっしょん", "passion"}
	for _, e := range keywords {
		if strings.Contains(strings.ToLower(text), e) {
			return true
		}
	}
	return false
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
				if !strings.Contains(botId, ev.User) {
					isMentioned := strings.Contains(ev.Text, botId)

					if IncludesPassion(ev.Text) {
						sendPassion(rtm, ev)
					} else if isMentioned && strings.Contains(ev.Text, "の画像") {
						sendGoogleImage(rtm, ev) // image search
					} else if strings.Contains(ev.Text, "申し訳ない") {
						sendImage(rtm, ev, "hakase.jpg", "本当に申し訳ない")
					} else if strings.Contains(ev.Text, "すみません") {
						sendImage(rtm, ev, "guusei.jpg", "ぐう聖すぎるほんとすみません")
					} else if strings.Contains(ev.Text, string(keyword2)) {
						sendAA(rtm, ev)
					} else if isMentioned && strings.Contains(ev.Text, "deleteAllMsgInChannel") {
						deleteAllMsg(api, ev) // clean up
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
