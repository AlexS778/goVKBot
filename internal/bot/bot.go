package bot

import (
	"encoding/json"
	"fmt"
	"goVkBot/internal/models"
	"goVkBot/internal/utils"
	"io"
	"log"
	"net/http"
	"net/url"
	"strconv"
)

type Bot struct {
	AccessToken   string
	GroupId       string
	ServerUrl     string
	LastTimeStamp string
}

// GetLongPollServer retrieves the long poll server information for a VK group.
// It makes a GET request to the VK API to obtain the server URL, key, and timestamp.
// The function takes a VK access token and group ID as parameters.
// It returns the constructed long poll URL and the timestamp for subsequent requests.
// In case of an error, the error is logged and an empty string is returned for the URL and timestamp.
func (b *Bot) GetLongPollServer() {
	// Construct the API URL
	serverUrl := fmt.Sprintf("https://api.vk.com/method/groups.getLongPollServer?access_token=%s&v=5.131&group_id=%s", b.AccessToken, b.GroupId)

	// Send a GET request to the VK API
	resp, err := http.Get(serverUrl)
	if err != nil {
		log.Println(err)
		return
	}
	defer resp.Body.Close()

	// Read the response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
		return
	}

	// Parse the response JSON
	var longPollServerCredentials struct {
		Response struct {
			Key    string `json:"key"`
			Server string `json:"server"`
			Ts     string `json:"ts"`
		} `json:"response"`
	}
	err = json.Unmarshal(body, &longPollServerCredentials)
	if err != nil {
		log.Println(err)
		return
	}

	// Construct the long poll URL
	longPollURL := fmt.Sprintf("%s?act=a_check&key=%s&ts=%s&wait=25", longPollServerCredentials.Response.Server, longPollServerCredentials.Response.Key, `%s`)

	b.ServerUrl = longPollURL
	b.LastTimeStamp = longPollServerCredentials.Response.Ts
}

// SendMessageToServer sends a message and an optional keyboard to the server using the VK API.
//
// Parameters:
//   - message: A string representing the message to be sent.
//   - response: A struct containing the server response data.
//   - keyboard: A struct representing the keyboard to be sent along with the message (optional).
//
// Note:
//   - The method prepares the necessary parameters and constructs the API URL to send the message
//     and keyboard (if provided) to the VK server.
//   - The method performs URL encoding for the message and keyboard parameters.
//   - If a keyboard is provided, it is marshaled to JSON format before being URL encoded.
//   - The user ID is extracted from the server response to determine the recipient of the message.
//   - A random ID is generated for each message sent.
//   - The VK API URL is constructed with the appropriate parameters, including access token, user ID,
//     random ID, message, and keyboard.
//   - If no keyboard is provided (empty string), the API URL is constructed without the keyboard parameter.
//   - The MakePostRequestWithUrl function is used to send the POST request to the VK API endpoint.
//   - Error handling is performed for parameter marshaling, URL encoding, and the POST request.
//     If an error occurs, an appropriate error message is logged.
func (b *Bot) SendMessageToServer(message string, response models.ServerResponse, keyboard models.Keyboard) {
	payload, err := json.Marshal(keyboard)
	if err != nil {
		log.Println("error marshaling the keyboard:", err)
		return
	}

	userId := strconv.Itoa(response.Updates[0].Object.Message.FromID)
	if userId == "0" {
		userId = strconv.Itoa(response.Updates[0].Object.PeerID)
	}
	randomId := utils.GetRandomInt32()
	encodedMessage := url.QueryEscape(message)
	encodedKeyboard := url.QueryEscape(string(payload))
	serverUrl := fmt.Sprintf("https://api.vk.com/method/messages.send?user_id=%s&random_id=%s&keyboard=%s&message=%s&access_token=%s&v=5.131", userId, randomId, encodedKeyboard, encodedMessage, b.AccessToken)
	//if there is no keyboard
	if encodedKeyboard == "" {
		serverUrl = fmt.Sprintf("https://api.vk.com/method/messages.send?user_id=%s&random_id=%s&message=%s&access_token=%s&v=5.131", userId, randomId, encodedMessage, b.AccessToken)
	}
	utils.MakePostRequestWithUrl(serverUrl)
}

// EditLastMessage edits the last sent message with the provided message content, keyboard, and conversation message ID.
//
// Parameters:
//   - message: A string representing the updated message content.
//   - response: A struct containing the server response data.
//   - cmId: A string representing the conversation message ID of the message to be edited.
//   - keyboard: A struct representing the updated keyboard (optional).
//
// Note:
//   - The method prepares the necessary parameters and constructs the API URL to edit the last sent message.
//   - The method performs URL encoding for the message and keyboard parameters.
//   - If a keyboard is provided, it is marshaled to JSON format before being URL encoded.
//   - The peer ID is extracted from the server response to identify the conversation.
//   - The VK API URL is constructed with the appropriate parameters, including peer ID,
//     message content, conversation message ID, keyboard, and access token.
//   - If no keyboard is provided (empty string), the API URL is constructed without the keyboard parameter.
//   - The MakePostRequestWithUrl function is used to send the POST request to the VK API endpoint.
//   - Error handling is performed for parameter marshaling, URL encoding, and the POST request.
//     If an error occurs, an appropriate error message is logged.
func (b *Bot) EditLastMessage(message string, response models.ServerResponse, cmId string, keyboard models.Keyboard) {
	payload, err := json.Marshal(keyboard)
	if err != nil {
		log.Println("error marshaling the keyboard:", err)
		return
	}
	encodedMessage := url.QueryEscape(message)
	encodedKeyboard := url.QueryEscape(string(payload))
	peerId := strconv.Itoa(response.Updates[0].Object.PeerID)
	serverUrl := fmt.Sprintf("https://api.vk.com/method/messages.edit?peer_id=%s&message=%s&conversation_message_id=%s&keyboard=%s&access_token=%s&v=5.131", peerId, encodedMessage, cmId, encodedKeyboard, b.AccessToken)
	//if there is no keyboard
	if encodedKeyboard == "" {
		serverUrl = fmt.Sprintf("https://api.vk.com/method/messages.edit?peer_id=%s&message=%s&conversation_message_id=%s&access_token=%s&v=5.131", peerId, encodedMessage, cmId, b.AccessToken)
	}
	utils.MakePostRequestWithUrl(serverUrl)
}

// HandleButtonCallback handles the callback event triggered by a button click.
//
// Parameters:
//   - eventData: A struct containing the event answer data received from the button callback.
//   - response: A struct containing the server response data.
//
// Note:
//   - The method prepares the necessary parameters and constructs the API URL to handle the button callback.
//   - The eventData parameter is marshaled to JSON format and URL encoded.
//   - The event ID and peer ID are extracted from the server response.
//   - The VK API URL is constructed with the appropriate parameters, including event ID, user ID (peer ID),
//     encoded event data, and access token.
//   - The MakePostRequestWithUrl function is used to send the POST request to the VK API endpoint.
//   - Error handling is performed for parameter marshaling and the POST request.
//     If an error occurs, an appropriate error message is logged.
func (b *Bot) HandleButtonCallback(eventData models.EventAnswer, response models.ServerResponse) {
	payload, err := json.Marshal(eventData)
	if err != nil {
		log.Println("error marshaling the keyboard:", err)
		return
	}
	encodedEventData := url.QueryEscape(string(payload))
	eventId := response.Updates[0].Object.EventID
	peerID := strconv.Itoa(response.Updates[0].Object.PeerID)
	serverUrl := fmt.Sprintf("https://api.vk.com/method/messages.sendMessageEventAnswer?event_id=%s&user_id=%s&peer_id=%s&event_data=%s&access_token=%s&v=5.131", eventId, peerID, peerID, encodedEventData, b.AccessToken)
	utils.MakePostRequestWithUrl(serverUrl)
}
