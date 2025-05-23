package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func makeRequest(url string) (int, string, error) {
	resp, err := http.Get(url)
	if err != nil {
		return 0, "", fmt.Errorf("Error making GET request: %w", err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return resp.StatusCode, "", fmt.Errorf("Error reading response body: %w", err)
	}

	return resp.StatusCode, string(body), nil
}

func main() {
	statusCode, body, err := makeRequest("https://jsonplaceholder.typicode.com/todos/1")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Status Code:", statusCode)
	fmt.Println("Response Body:", body)
}
