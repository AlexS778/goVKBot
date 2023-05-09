package utils

import (
	"encoding/json"
	"fmt"
	"goVkBot/internal/models"
	"io"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"time"
)

// GetRandomInt32 generates a random int32 number and returns it as a string.
// The function seeds the random number generator with the current time to ensure randomness.
// It uses the built-in `rand` package to generate the random number.
// The generated random number is then converted to a string using `strconv.Itoa`.
// The function does not take any parameters.
// It returns a string representation of the generated random int32 number.
func GetRandomInt32() string {
	// Seed the random number generator with the current time
	rand.NewSource(time.Now().UnixNano())

	// Generate a random int32 number
	randomNumber := rand.Int31()

	// Convert the random number to a string
	return strconv.Itoa(int(randomNumber))
}

// CreateButton creates a button with the specified label, link, color, type, and payload.
// It returns a models.Button struct representing the created button.
//
// Parameters:
//   - label: The label or text displayed on the button.
//   - link: The URL link associated with the button (optional, can be an empty string).
//   - color: The color of the button (optional, can be an empty string).
//   - typeOfButton: The type of the button, such as "text" or "open_link".
//   - payload: The payload associated with the button (optional, can be an empty string).
//
// Returns:
//   - models.Button: The created button with the specified properties.
func CreateButton(label string, link string, color string, typeOfButton string, payload string) models.Button {
	button := models.Button{}
	button.Action.Label = label
	button.Action.Link = link
	button.Color = color
	button.Action.Type = typeOfButton
	button.Action.Payload = payload
	return button
}

// GetWeatherInfo retrieves weather information for the specified city.
// It makes a request to the weather service API and returns the current temperature as a string.
//
// Parameters:
//   - city: The name of the city for which weather information is requested.
//
// Returns:
//   - string: The current temperature in the specified city as a string.
//     If an error occurs during the retrieval or parsing of weather data, an empty string is returned.
//
// Supported Cities:
//   - Moscow: Retrieves weather information for Moscow, Russia.
//   - London: Retrieves weather information for London, United Kingdom.
//
// Note:
//   - The weather service API used in this function may have limitations or require proper authentication.
//     Ensure that you have the necessary permissions and provide valid API endpoints for other cities if needed.
func GetWeatherInfo(city string) string {
	cityUrl := ""
	if city == "Moscow" {
		cityUrl = "https://api.open-meteo.com/v1/forecast?latitude=55.75&longitude=37.62&hourly=temperature_2m&current_weather=true&forecast_days=1&timezone=Europe%2FMoscow"
	}
	if city == "London" {
		cityUrl = "https://api.open-meteo.com/v1/forecast?latitude=51.51&longitude=-0.13&hourly=temperature_2m&current_weather=true"
	}
	response, err := http.Get(cityUrl)
	if err != nil {
		log.Println("error making request to weather service:", err)
		return ""
	}
	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		log.Println("error reading response body", err)
		return ""
	}

	weather := models.Weather{}
	err = json.Unmarshal(body, &weather)
	if err != nil {
		log.Println("error unmarshalling body", err)
		return ""
	}
	return strconv.FormatFloat(weather.CurrentWeather.Temperature, 'f', -1, 64)
}

// GetRandomCat retrieves a URL of a random cat image from the Cataas service.
//
// Returns:
//   - A string containing the URL of a random cat image.
//   - An empty string if an error occurs during the HTTP request or response handling.
//
// Note:
//   - The function makes a GET request to the Cataas service to retrieve a random cat image.
//   - The URL of the cat image is extracted from the response body.
//   - Error handling is performed for the HTTP request creation and response handling.
//     If an error occurs, an error message is logged, and an empty string is returned.
func GetRandomCat() string {
	catUrl := "https://cataas.com/cat?json=true"
	response, err := http.Get(catUrl)
	if err != nil {
		log.Println(err)
		return ""
	}
	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		log.Println("error reading response body", err)
		return ""
	}

	catData := models.Cat{}
	err = json.Unmarshal(body, &catData)
	if err != nil {
		log.Println("error unmarshalling body", err)
		return ""
	}
	return "https://cataas.com/" + catData.URL
}

// MakePostRequestWithUrl makes a POST request to the specified URL without sending any request body.
// It sends an empty form body to the URL and prints the response body and status to the standard output.
//
// Parameters:
//   - url: The URL to which the POST request is made.
//
// Note:
//   - This function does not send any request body. It only sends an empty form body.
//   - The response body and status are printed to the standard output for debugging or informational purposes.
//   - Error handling is performed for the HTTP request creation, but not for reading the response body.
//     Make sure to handle errors appropriately in a production environment.
func MakePostRequestWithUrl(url string) {
	resp, err := http.PostForm(url, nil)
	if err != nil {
		fmt.Println("Error creating request:", err)
		return
	}

	body, _ := io.ReadAll(resp.Body)
	fmt.Println(string(body))

	defer resp.Body.Close()

	fmt.Println("Response status:", resp.Status)
}
