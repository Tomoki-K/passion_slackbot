package controllers

import (
	"encoding/base64"
	"io/ioutil"
	"log"
	"math/rand"
	"strings"
	"time"

	"github.com/Tomoki-K/passion_slackbot/bobuneMaker"
	"github.com/Tomoki-K/passion_slackbot/image"
	"github.com/Tomoki-K/passion_slackbot/models"
	"github.com/nlopes/slack"
)

type MsgController struct {
	Rtm    *slack.RTM
	Ev     *slack.MessageEvent
	Client *slack.Client
	BotId  string
}

var passionMsg = "パッションが足りません。"
var rareMsg = "温かいし止まらない。"
var srareMsg = ":v::gunma::v:"
var aa = "" +
	"44CA44CA44CAIF8g44CA4oipDQrjgIDjgIAo44C" +
	"A776f4oiA776fKeW9oeOAgOOBiuOBo+OBseOBhC" +
	"HjgYrjgaPjgbHjgYQhDQrjgIDjgIAo44CAIOKKg" +
	"uW9oQ0K44CAIOOAgHzjgIDjgIDjgIB8DQrjgIAg" +
	"44CA44GXIOKMku+8qg=="

func NewMsgController(rtm *slack.RTM, ev *slack.MessageEvent, api *slack.Client, id string) MsgController {
	return MsgController{Rtm: rtm, Ev: ev, Client: api, BotId: id}
}

func decodeAA(encAA string) string {
	b, err := base64.StdEncoding.DecodeString(encAA)
	if err != nil {
		return passionMsg // fallback to default if err occures
	}
	return "\n" + string(b)
}

func hasPassionedToday(user string, hist []models.History) bool {
	today := time.Now().Format("2006-01-02")
	for _, h := range hist {
		if h.UserID == user && h.LastPassion.Format("2006-01-02") == today {
			return true
		}
	}
	return false
}

func (c MsgController) SendPassion(api *slack.Client, history []models.History) []models.History {
	mentionTag := "<@" + c.Ev.User + "> "

	if hasPassionedToday(c.Ev.User, history) {
		sender, err := api.GetUserInfo(c.Ev.User)
		filename, err := bobune.CreateBobuneImg(sender)
		if err != nil {
			log.Print(err)
		}
		var fileParams = slack.FileUploadParameters{
			File:     "assets/out/" + filename,
			Filetype: "image",
			Filename: "",
			Channels: []string{c.Ev.Channel},
		}
		c.Rtm.UploadFile(fileParams)
		return history

	} else {
		text := mentionTag + passionMsg // default message
		rand.Seed(time.Now().UTC().UnixNano())
		randInt := rand.Intn(100)
		if randInt < 1 {
			text = mentionTag + srareMsg // 1% chance of rare message
		} else if randInt < 5 {
			text = mentionTag + rareMsg // 5% chance of rare message
		}
		c.Rtm.SendMessage(c.Rtm.NewOutgoingMessage(text, c.Ev.Channel))

		// write history
		var newHistory []models.History
		for _, h := range history {
			if h.UserID == c.Ev.User {
				h.LastPassion = time.Now()
			}
			newHistory = append(newHistory, h)
		}
		return newHistory
	}
}

func (c MsgController) SendGoogleImage() {
	mentionTag := "<@" + c.Ev.User + "> "
	searchWord := strings.Replace(c.Ev.Text, "の画像", "", -1)
	searchWord = strings.Replace(searchWord, c.BotId, "", -1)
	imgUrl, err := image.ImageSearch(searchWord)
	if err != nil || len(strings.TrimSpace(searchWord)) < 1 {
		imgUrl = "invalid search. (Usage: '@passion_bot 〇〇の画像')"
		log.Print(err)
	}
	c.Rtm.SendMessage(c.Rtm.NewOutgoingMessage(mentionTag+imgUrl, c.Ev.Channel))
}

func (c MsgController) SendImage(filename string, title string) {
	var fileParams = slack.FileUploadParameters{
		File:     "assets/" + filename,
		Filetype: "image",
		Filename: title,
		Channels: []string{c.Ev.Channel},
	}
	c.Rtm.UploadFile(fileParams)
}

func (c MsgController) SendAA() {
	mentionTag := "<@" + c.Ev.User + "> "
	var text = mentionTag + decodeAA(aa)
	c.Rtm.SendMessage(c.Rtm.NewOutgoingMessage(text, c.Ev.Channel))
}

func (c MsgController) DeleteAllMsg() {
	var params = slack.GetConversationHistoryParameters{ChannelID: c.Ev.Channel, Limit: 100}
	hist, err := c.Client.GetConversationHistory(&params)
	if err != nil {
		log.Print(err)
	} else {
		for _, v := range hist.Messages {
			c.Client.DeleteMessage(c.Ev.Channel, v.Timestamp)
			log.Println(v.Text + " ===> deleted")
			time.Sleep(100 * time.Millisecond)
		}
		log.Println("done!")
	}
}

func (c MsgController) SendHelp() {
	buf, err := ioutil.ReadFile("features.txt")
	if err != nil {
		log.Print(err)
	}
	s := string(buf)
	c.Rtm.SendMessage(c.Rtm.NewOutgoingMessage(s, c.Ev.Channel))
}
