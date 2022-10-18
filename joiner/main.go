package main

import (
	"log"
	"net/http"
	"net/url"
	"os"
	"strconv"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/buger/jsonparser"
)

var telegramApi string = "https://api.telegram.org/bot" + os.Getenv("TELEGRAM_BOT_TOKEN") + "/deleteMessage"

func handler(request events.LambdaFunctionURLRequest) {
	payload := []byte(request.Body)
	_, err := jsonparser.GetInt(payload, "message", "new_chat_participant", "id")
	if err != nil {
		return
	}
	chat, err := jsonparser.GetInt(payload, "message", "chat", "id")
	if err != nil {
		log.Fatal(err)
	}
	messageID, err := jsonparser.GetInt(payload, "message", "message_id")
	if err != nil {
		log.Fatal(err)
	}
	// delete message
	_, err = http.PostForm(
		telegramApi,
		url.Values{
			"chat_id":    {strconv.Itoa(int(chat))},
			"message_id": {strconv.Itoa(int(messageID))},
		})

	if err != nil {
		log.Fatal(err)
	}

}

func main() {
	lambda.Start(handler)
}
