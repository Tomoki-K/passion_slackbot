package main

import (
	"log"
	"os"
	"strings"

	"github.com/nlopes/slack"
)

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
				isPassion, err := IncludesPassion(ev.Text)
				if err != nil {
					log.Fatal(err)
				}
				if isPassion {
					rtm.SendMessage(rtm.NewOutgoingMessage("パッションが足りません。", ev.Channel))
				}

			case *slack.InvalidAuthEvent:
				log.Fatal("Invalid credentials")
			}
		}
	}
}

func main() {
	api := slack.New(os.Getenv("SLACK_PASSION_KEY"))
	os.Exit(run(api))
}
