package main

import (
	"fmt"
	"goVkBot/internal/bot"
	"goVkBot/internal/models"
	"goVkBot/internal/server"
	"goVkBot/internal/utils"
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file:", err)
	}
	token := os.Getenv("TOKEN")
	groupId := os.Getenv("GROUPID")

	// populate the bot with data
	myBot := bot.Bot{AccessToken: token, GroupId: groupId}

	// receieve address of LongPollServer
	myBot.GetLongPollServer()

	// create a channel to listen for responses from LongPollServer
	responseChan := make(chan models.ServerResponse, 1)

	// start a goroutine that listens to the LongPollServer continiously
	go server.ListenForResponses(myBot, responseChan)

	// map to contain last messages sent by bot
	lastMessageId := map[int]int{}

	// handle the responses from the LongPollServer accordingly
	for response := range responseChan {
		if len(response.Updates) > 0 {
			userMessage := response.Updates[0].Object.Message.Text
			payload := response.Updates[0].Object.Payload
			if userMessage == "Начать" {
				weatherButton := utils.CreateButton("Получить погоду", "", "primary", "text", "")
				googleButton := utils.CreateButton("Go to google.com", "https://google.com", "", "open_link", "")
				catsButton := utils.CreateButton("Получить фото кота!", "", "", "text", "")
				bookTable := utils.CreateButton("Забронировать столик", "", "primary", "text", "")
				keyboard := models.Keyboard{Inline: false, Buttons: [][]models.Button{{weatherButton}, {googleButton}, {catsButton}, {bookTable}}}
				myBot.SendMessageToServer("Привет! Этот бот был сделан для VK \n Выбери что-то из кнопок снизу:", response, keyboard)
			}
			if userMessage == "Получить погоду" {
				temperature := utils.GetWeatherInfo("Moscow")
				message := fmt.Sprintf("Погода в Москве: %s \u2103", temperature)
				moscowButton := utils.CreateButton("Москва", "", "", "callback", "{\"button\": \"moscow\"}")
				londonButton := utils.CreateButton("London", "", "", "callback", "{\"button\": \"london\"}")
				keyboard := models.Keyboard{Inline: true, Buttons: [][]models.Button{{moscowButton}, {londonButton}}}
				myBot.SendMessageToServer(message, response, keyboard)
				lastMessageId[response.Updates[0].Object.Message.PeerID] = response.Updates[0].Object.Message.ConversationMessageID + 1
			}
			if userMessage == "Получить фото кота!" {
				catPicture := utils.GetRandomCat()
				myBot.SendMessageToServer(catPicture, response, models.Keyboard{})
			}
			if fmt.Sprintf("%s", payload) == "{moscow}" {
				temperature := utils.GetWeatherInfo("Moscow")
				message := fmt.Sprintf("Погода в Москве: %s \u2103", temperature)
				cmId := strconv.Itoa(lastMessageId[response.Updates[0].Object.PeerID])
				moscowButton := utils.CreateButton("Москва", "", "", "callback", "{\"button\": \"moscow\"}")
				londonButton := utils.CreateButton("London", "", "", "callback", "{\"button\": \"london\"}")
				keyboard := models.Keyboard{Inline: true, Buttons: [][]models.Button{{moscowButton}, {londonButton}}}
				myBot.EditLastMessage(message, response, cmId, keyboard)
			}
			if fmt.Sprintf("%s", payload) == "{london}" {
				temperature := utils.GetWeatherInfo("London")
				message := fmt.Sprintf("Погода в Лондоне: %s \u2103", temperature)
				cmId := strconv.Itoa(lastMessageId[response.Updates[0].Object.PeerID])
				moscowButton := utils.CreateButton("Москва", "", "", "callback", "{\"button\": \"moscow\"}")
				londonButton := utils.CreateButton("London", "", "", "callback", "{\"button\": \"london\"}")
				keyboard := models.Keyboard{Inline: true, Buttons: [][]models.Button{{moscowButton}, {londonButton}}}
				myBot.EditLastMessage(message, response, cmId, keyboard)
			}
			if userMessage == "Забронировать столик" {
				time1600 := utils.CreateButton("16:00", "", "", "callback", "{\"button\": \"time\"}")
				time1700 := utils.CreateButton("17:00", "", "", "callback", "{\"button\": \"time\"}")
				time1800 := utils.CreateButton("18:00", "", "", "callback", "{\"button\": \"time\"}")
				time1900 := utils.CreateButton("19:00", "", "", "callback", "{\"button\": \"time\"}")
				keyboard := models.Keyboard{Inline: true, Buttons: [][]models.Button{{time1600}, {time1700}, {time1800}, {time1900}}}
				myBot.SendMessageToServer("Выберите время:", response, keyboard)
			}
			if fmt.Sprintf("%s", payload) == "{time}" {
				eventData := models.EventAnswer{Type: "show_snackbar", Text: "Время подтверждено!"}
				myBot.HandleButtonCallback(eventData, response)
				yesButton := utils.CreateButton("Да", "", "positive", "callback", "{\"button\": \"confirm\"}")
				noButton := utils.CreateButton("Нет", "", "negative", "callback", "{\"button\": \"back\"}")
				keyboard := models.Keyboard{Inline: true, Buttons: [][]models.Button{{yesButton}, {noButton}}}
				myBot.SendMessageToServer("Подтвердить бронь?", response, keyboard)
			}
			if fmt.Sprintf("%s", payload) == "{confirm}" {
				eventData := models.EventAnswer{Type: "show_snackbar", Text: "Ваша заявка принята! \nМенеджер свяжется с вами в течение часа для потверждения брони."}
				myBot.HandleButtonCallback(eventData, response)
				myBot.SendMessageToServer("Вы сделали заявку, ождидайте звонка менеджера.", response, models.Keyboard{})

			}
			if fmt.Sprintf("%s", payload) == "{back}" {
				eventData := models.EventAnswer{Type: "show_snackbar", Text: "Вы вернулись назад."}
				myBot.HandleButtonCallback(eventData, response)
				weatherButton := utils.CreateButton("Получить погоду", "", "", "text", "")
				googleButton := utils.CreateButton("Go to google.com", "", "https://google.com", "open_link", "")
				catsButton := utils.CreateButton("Получить фото кота!", "", "", "text", "")
				bookTable := utils.CreateButton("Забронировать столик", "", "", "text", "")
				keyboard := models.Keyboard{Inline: false, Buttons: [][]models.Button{{weatherButton}, {googleButton}, {catsButton}, {bookTable}}}
				myBot.SendMessageToServer("Привет! Этот бот был сделан для VK \n Выбери что-то из кнопок снизу:", response, keyboard)
			}
		}
	}
}
