package main

import (
	"encoding/base64"
	"log"
	"os"
	"strings"

	"github.com/Tomoki-K/passion_slackbot/controllers"
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

				mc := controllers.NewMsgController(rtm, ev, api, "<@UA26NB6H0>")

				if !strings.Contains(mc.BotId, ev.User) {
					isMentioned := strings.Contains(ev.Text, mc.BotId)
					keyword2, _ := base64.StdEncoding.DecodeString("44GK44Gj44Gx44GE")

					if isMentioned {
						if strings.Contains(ev.Text, "の画像") {
							mc.SendGoogleImage() // image search

						} else if strings.Contains(ev.Text, "deleteAllMsgInChannel") {
							mc.DeleteAllMsg() // clean up
						} else if ev.Text == mc.BotId+" help" {
							mc.SendHelp()
						} else {
							mc.Rtm.SendMessage(mc.Rtm.NewOutgoingMessage("`@passion_bot help` for help", mc.Ev.Channel))
						}
					} else if IncludesPassion(ev.Text) {
						mc.SendPassion()

					} else if strings.Contains(ev.Text, "申し訳ない") {
						mc.SendImage("hakase.jpg", "本当に申し訳ない")

					} else if strings.Contains(ev.Text, "すみません") {
						mc.SendImage("guusei.jpg", "ぐう聖すぎるほんとすみません")

					} else if strings.Contains(ev.Text, string(keyword2)) {
						mc.SendAA()
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
