package models

import "time"

// Longpoll server response
type ServerResponse struct {
	Ts      string `json:"ts"`
	Updates []struct {
		GroupID int    `json:"group_id"`
		Type    string `json:"type"`
		EventID string `json:"event_id"`
		V       string `json:"v"`
		Object  struct {
			Message struct {
				Date                  int           `json:"date"`
				FromID                int           `json:"from_id"`
				ID                    int           `json:"id"`
				Out                   int           `json:"out"`
				Attachments           []interface{} `json:"attachments"`
				ConversationMessageID int           `json:"conversation_message_id"`
				FwdMessages           []interface{} `json:"fwd_messages"`
				Important             bool          `json:"important"`
				IsHidden              bool          `json:"is_hidden"`
				PeerID                int           `json:"peer_id"`
				RandomID              int           `json:"random_id"`
				Text                  string        `json:"text"`
			} `json:"message"`
			ClientInfo struct {
				ButtonActions  []string `json:"button_actions"`
				Keyboard       bool     `json:"keyboard"`
				InlineKeyboard bool     `json:"inline_keyboard"`
				Carousel       bool     `json:"carousel"`
				LangID         int      `json:"lang_id"`
			} `json:"client_info"`
			Payload struct {
				Button string `json:"button"`
			} `json:"payload"`
			PeerID                int    `json:"peer_id"`
			ConversationMessageID int    `json:"conversation_message_id"`
			EventID               string `json:"event_id"`
		} `json:"object"`
	} `json:"updates"`
}

// Keyboard struct that is being sent with message
type Keyboard struct {
	Inline  bool       `json:"inline,omitempty"`
	Buttons [][]Button `json:"buttons,omitempty"`
}

// Button struct that is being sent with Keyboard
type Button struct {
	Action struct {
		Type    string `json:"type"`
		Link    string `json:"link,omitempty"`
		Label   string `json:"label"`
		Payload string `json:"payload,omitempty"`
	} `json:"action"`
	Color string `json:"color,omitempty"`
}

// Weather struct that represents json response from weather service
type Weather struct {
	Latitude             float64 `json:"latitude"`
	Longitude            float64 `json:"longitude"`
	GenerationtimeMs     float64 `json:"generationtime_ms"`
	UtcOffsetSeconds     int     `json:"utc_offset_seconds"`
	Timezone             string  `json:"timezone"`
	TimezoneAbbreviation string  `json:"timezone_abbreviation"`
	Elevation            float64 `json:"elevation"`
	CurrentWeather       struct {
		Temperature   float64 `json:"temperature"`
		Windspeed     float64 `json:"windspeed"`
		Winddirection float64 `json:"winddirection"`
		Weathercode   int     `json:"weathercode"`
		IsDay         int     `json:"is_day"`
		Time          string  `json:"time"`
	} `json:"current_weather"`
	HourlyUnits struct {
		Time          string `json:"time"`
		Temperature2M string `json:"temperature_2m"`
	} `json:"hourly_units"`
	Hourly struct {
		Time          []string  `json:"time"`
		Temperature2M []float64 `json:"temperature_2m"`
	} `json:"hourly"`
}

// Cat struct that represents json response from cats image service
type Cat struct {
	Tags      []string  `json:"tags"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
	Validated bool      `json:"validated"`
	Owner     string    `json:"owner"`
	File      string    `json:"file"`
	Mimetype  string    `json:"mimetype"`
	Size      int       `json:"size"`
	ID        string    `json:"_id"`
	URL       string    `json:"url"`
}

// EventAnswer struct that is being sent to response callback
type EventAnswer struct {
	Type string `json:"type"`
	Text string `json:"text"`
}
