package main

import (
	"log"
	"math/rand"
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
					log.Print(err)
					return 1
				}
				if isPassion {
					// mention sender
					text := "<@" + ev.User + "> パッションが足りません。"
					if rand.Intn(100) < 5 {
						text = "<@" + ev.User + "> 温かいし止まらない。"
					}
					rtm.SendMessage(rtm.NewOutgoingMessage(text, ev.Channel))
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
