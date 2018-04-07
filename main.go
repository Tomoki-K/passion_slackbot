package main

import (
	"log"
	"os"
	"strings"

	"github.com/nlopes/slack"
)

func includes_passion(text string) bool {
	keywords := [...]string{"パッション", "ぱっしょん", "passion", "Passion"}
	for _, e := range keywords {
		if strings.Contains(text, e) {
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
				if includes_passion(ev.Text) {
					rtm.SendMessage(rtm.NewOutgoingMessage("パッションが足りません。", ev.Channel))
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
