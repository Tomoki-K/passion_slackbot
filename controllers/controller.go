package controllers

import (
	"encoding/base64"
	"log"
	"math/rand"
	"strings"
	"time"

	"github.com/Tomoki-K/passion_slackbot/image"
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

type MsgController struct {
	Rtm    *slack.RTM
	Ev     *slack.MessageEvent
	Client *slack.Client
	BotId  string
}

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

func (c MsgController) SendPassion() {
	mentionTag := "<@" + c.Ev.User + "> "
	text := mentionTag + passionMsg // default message
	rand.Seed(time.Now().UTC().UnixNano())
	if rand.Intn(100) < 5 {
		text = mentionTag + rareMsg // 5% chance of rare message
	}
	c.Rtm.SendMessage(c.Rtm.NewOutgoingMessage(text, c.Ev.Channel))
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
		Filename: "本当に申し訳ない",
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
