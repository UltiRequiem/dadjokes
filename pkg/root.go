package dadjokes

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

const API_URL = "https://icanhazdadjoke.com"

type Joke struct {
	Joke string `json:"joke"`
	Id   string `json:"id"`
}

func GetJoke() (Joke, error) {
	req, errorRequesting := http.NewRequest("GET", API_URL, nil)

	if errorRequesting != nil {
		return Joke{}, errorRequesting
	}

	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")

	client := &http.Client{}

	response, errorFetching := client.Do(req)

	if errorFetching != nil {
		return Joke{}, errorFetching
	}

	defer response.Body.Close()

	body, errorReading := ioutil.ReadAll(response.Body)

	if errorReading != nil {
		return Joke{}, errorReading
	}

	var joke Joke

	errorUnmarshalling := json.Unmarshal(body, &joke)

	if errorUnmarshalling != nil {
		return Joke{}, errorUnmarshalling
	}

	return joke, nil
}
