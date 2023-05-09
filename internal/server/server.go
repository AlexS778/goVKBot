package server

import (
	"encoding/json"
	"fmt"
	"goVkBot/internal/bot"
	"goVkBot/internal/models"
	"io"
	"log"
	"net/http"
	"time"
)

// ListenForResponses listens for updates from a VK long poll server and sends the updates to a channel for processing.
// It takes the server URL, the timestamp of the last update received, and a response channel as parameters.
// The function makes periodic GET requests to the server to check for new updates.
// When an update is received, it is unmarshaled into a `models.ServerResponse` struct.
// The updated timestamp is extracted from the response and used for the next request.
// The received response is sent to the provided response channel for further processing.
func ListenForResponses(b bot.Bot, responseChan chan<- models.ServerResponse) {
	baseUrl := b.ServerUrl
	for {
		resp, err := http.Get(fmt.Sprintf(baseUrl, b.LastTimeStamp))
		if err != nil {
			log.Println("Error making the request:", err)
			continue
		}

		body, err := io.ReadAll(resp.Body)
		if err != nil {
			log.Println("Error reading the response:", err)
			continue
		}

		log.Println(string(body))

		updateResponseData := models.ServerResponse{}
		json.Unmarshal(body, &updateResponseData)
		b.LastTimeStamp = updateResponseData.Ts

		resp.Body.Close()

		responseChan <- updateResponseData

		time.Sleep(time.Second * 1)
	}
}
