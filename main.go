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

func parseHttpResponse(response *http.Response) []byte {
	responseData, readError := ioutil.ReadAll(response.Body)

	if readError != nil {
		fmt.Println("Error reading response.")
		os.Exit(1)
	}

	return responseData
}

func parseBytesToResponseObject(bytes []byte) Response {
	var responseObject Response

	json.Unmarshal(bytes, &responseObject)
	return responseObject
}

func main() {
	client := &http.Client{}

	resp, err := client.Do(fetch("https://icanhazdadjoke.com"))

	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}

	defer resp.Body.Close()

	responseObject := parseBytesToResponseObject(parseHttpResponse(resp))

	fmt.Println(" " + responseObject.Joke)
}
