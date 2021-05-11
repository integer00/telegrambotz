package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"time"

	log "github.com/sirupsen/logrus"
)

var TelegramToken = os.Getenv("TELEGRAM_TOKEN")
var TelegramApiAddress = "https://api.telegram.org/bot"

type Chat struct {
	Id int `json:"id"`
}
type Message struct {
	Text string `json:"text"`
	Chat Chat   `json:"chat"`
}
type Update struct {
	UpdateId int     `json:"update_id"`
	Message  Message `json:"message"`
}

type Updates struct {
	Ok     bool `json:"ok"`
	Result []Update
}

func parseTelegramUpdates(r *http.Response) (*Updates, error) {
	var updates Updates
	if err := json.NewDecoder(r.Body).Decode(&updates); err != nil {
		log.Errorf("could't decode getUpdates %s", err.Error())
		return nil, err
	}

	return &updates, nil
}
func parseTelegramRequest(r *http.Request) (*Update, error) {
	var update Update
	if err := json.NewDecoder(r.Body).Decode(&update); err != nil {
		log.Errorf("could not decode incoming update %s", err.Error())
		return nil, err
	}

	return &update, nil
}

func getUpdates() *Updates {

	var getUpdatesAddress = TelegramApiAddress + TelegramToken + "/getUpdates"

	response, err := http.Get(getUpdatesAddress)
	if err != nil {
		log.Errorf("failed to get updates, %s", err.Error())
	}
	updates, err := parseTelegramUpdates(response)
	if err != nil {
		log.Error(err.Error())
	}

	return updates

}

func listenUpdates() {
	var offset = 0

	log.Info("listening for updates...")
	for {
		updates := getUpdates()
		for _, update := range updates.Result {
			if update.UpdateId > offset {
				offset = update.UpdateId
				log.Info("got new update, replying...")
				sendToTelegram(update.Message.Chat.Id, "reply")
			}

		}
		time.Sleep(15 * time.Second)
	}
}

func sendToTelegram(chatId int, text string) (string, error) {

	var sendAddress = TelegramApiAddress + TelegramToken + "/sendMessage"

	log.Infof("sending %s to chat_id: %d \n", text, chatId)
	log.Debugf("sending to: %s", sendAddress)
	response, err := http.PostForm(
		sendAddress,
		url.Values{
			"chat_id": {strconv.Itoa(chatId)},
			"text":    {text},
		})
	if err != nil {
		panic(err.Error())
	}
	defer response.Body.Close()

	var bodyBytes, errRead = ioutil.ReadAll(response.Body)
	if errRead != nil {
		panic(err.Error())
	}
	bodyString := string(bodyBytes)
	log.Infof("response: %s", bodyString)

	return bodyString, nil

}

func handleTelegramHook(w http.ResponseWriter, r *http.Request) {
	var update, err = parseTelegramRequest(r)
	if err != nil {
		log.Errorf("error parsing update: %s", err.Error())
		return
	}

	sendToTelegram(update.Message.Chat.Id, "hello")
	if err != nil {
		log.Errorf("error while sending to telegram")
		return
	}
}

func localTest() {
	log.Infof("starting local test polling\n")
	log.Debugf("token is: %s", os.Getenv("TELEGRAM_TOKEN"))
	http.HandleFunc("/", handleTelegramHook)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func main() {
	log.SetLevel(log.DebugLevel)
	go listenUpdates()
	localTest()
}
