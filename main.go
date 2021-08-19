package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

type Response struct {
	Joke string `json:"joke"`
}

func fetch(url string) *http.Request {
	req, err := http.NewRequest("GET", url, nil)

	if err != nil {
		fmt.Println("Probably you don't have internet.")
		os.Exit(1)
	}

	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")
	return req
}



func main() {
	client := &http.Client{}
	resp, err := client.Do(fetch("https://icanhazdadjoke.com"))

	if err != nil {
		fmt.Print(err.Error())
	}

	defer resp.Body.Close()

	bodyBytes, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		fmt.Print(err.Error())
	}

	var responseObject Response

	json.Unmarshal(bodyBytes, &responseObject)

	fmt.Println(" " + responseObject.Joke)
}
