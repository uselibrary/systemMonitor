package teleBot

import (
	"log"
	"net/http"
)

// function to send message to telegram bot
// variable: TOKEN, CHAT_ID, message
// return: True if success, False if failed
func SendMessage(token string, chatID string, message string) bool {
	// create url
	url := "https://api.telegram.org/bot" + token + "/sendMessage?chat_id=" + chatID + "&text=" + message

	// send request
	resp, err := http.Get(url)
	if err != nil {
		log.Println(err)
		return false
	}

	// check response
	if resp.StatusCode == 200 {
		return true
	} else {
		return false
	}
}

