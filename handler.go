package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strconv"
	"os"
)

const botName string = "@carRegionBot"

const startCommand string = "/start"

const startInfo string = "Now you can enter registation region code (ex. 01, 05, 18)."


// Pass token and sensible APIs through environment variables
const telegramApiBaseUrl string = "https://api.telegram.org/bot"
const telegramApiSendMessage string = "/sendMessage"
const telegramTokenEnv string = "TELEGRAM_BOT_TOKEN"

var telegramApi string = telegramApiBaseUrl + os.Getenv(telegramTokenEnv) + telegramApiSendMessage


// Update is a Telegram object that we receive every time an user interacts with the bot.
type Update struct {
	UpdateId int     `json:"update_id"`
	Message  Message `json:"message"`
}

// Implements the fmt.String interface to get the representation of an Update as a string.
func (u Update) String() string {
	return fmt.Sprintf("(update id: %d, message: %s)", u.UpdateId, u.Message)
}

// Message is a Telegram object that can be found in an update.
// Note that not all Update contains a Message. Update for an Inline Query doesn't.
type Message struct {
	Text     string   `json:"text"`
	Chat     Chat     `json:"chat"`
}

// A Chat indicates the conversation to which the Message belongs.
type Chat struct {
	Id int `json:"id"`
}

// Implements the fmt.String interface to get the representation of a Chat as a string.
func (c Chat) String() string {
	return fmt.Sprintf("(id: %d)", c.Id)
}

type Regions struct {
    Regions []Region `json:"data"`
}

type Region struct {
    Regioncode   string `json:"regioncode"`
	Offname 	 string `json:"offname"`
    Shortname    string `json:"shortname"`
}

/**
 * Handle Telegram webhook request
 */
func HandleTelegramWebHook(w http.ResponseWriter, r *http.Request) {

	// Parse incoming request
	var update, err = parseTelegramRequest(r)
	if err != nil {
		log.Printf("error parsing update, %s", err.Error())
		return
	}

	// Send start instrutions
	var text = botName + ": Invalid code";
	if (update.Message.Text == startCommand) {
		text = startInfo;
	} else if _, err := strconv.Atoi(update.Message.Text); err == nil {
		text = getRegionName(update.Message.Text)
	}
	// Send region name to Telegram
	var telegramResponseBody, errTelegram = sendTextToTelegramChat(update.Message.Chat.Id, text)
	if errTelegram != nil {
		log.Printf("got error %s from telegram, response body is %s", errTelegram.Error(), telegramResponseBody)
	} else {
		log.Printf("Message %s successfully distributed to chat id %d", text, update.Message.Chat.Id)
	}
}

/**
 * Handles incoming update from the Telegram web hook
 */
func parseTelegramRequest(r *http.Request) (*Update, error) {
	var update Update
	if err := json.NewDecoder(r.Body).Decode(&update); err != nil {
		log.Printf("could not decode incoming update %s", err.Error())
		return nil, err
	}
	if update.UpdateId == 0 {
		log.Printf("invalid update id, got update id = 0")
		return nil, errors.New("invalid update id of 0 indicates failure to parse incoming update")
	}
	return &update, nil
}

/**
 * Get region name by code
 */
func getRegionName(code string) string {

	// Open our jsonFile
	// Source http://basicdata.ru/api/json/fias/addrobj
	jsonFile, err := os.Open("regions.json")

	// if we os.Open returns an error then handle it
	if err != nil {
		fmt.Println(err)
	}

	// defer the closing of our jsonFile so that we can parse it later on
	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)

	var regions Regions

	json.Unmarshal(byteValue, &regions)
	
	for i := 0; i < len(regions.Regions); i++ {
		if (regions.Regions[i].Regioncode == code) {
			return fmt.Sprintf("%s (%s)", regions.Regions[i].Offname, regions.Regions[i].Shortname)
		}
	}

	return  botName + ": Сode does not exist."
}

/**
  * Sends a text message to the Telegram chat identified by its chat Id
  */
func sendTextToTelegramChat(chatId int, text string) (string, error) {

	log.Printf("Sending %s to chat_id: %d", text, chatId)
	response, err := http.PostForm(
		telegramApi,
		url.Values{
			"chat_id": {strconv.Itoa(chatId)},
			"text":    {text},
		})

	if err != nil {
		log.Printf("Error when posting text to the chat: %s", err.Error())
		return "", err
	}
	defer response.Body.Close()

	var bodyBytes, errRead = ioutil.ReadAll(response.Body)
	if errRead != nil {
		log.Printf("Error in parsing telegram answer %s", errRead.Error())
		return "", err
	}
	bodyString := string(bodyBytes)

	return bodyString, nil
}