package provider

import (
	"fmt"
	"log"
	"net/http"
)

func GetInfo() (*http.Response, error) {
	url := "https://api.football-data.org/v4/teams/81/matches?status=SCHEDULED&limit=4"
	apiToken := "bd2bea287209438496914236fb543609"
	tokenHeader := "X-Auth-Token"
	// 1. Create a new request object
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatal(err)
	}

	// 2. Add or set headers
	req.Header.Add(tokenHeader, apiToken)

	// 3. Send the request using a client
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	// Handle the response
	fmt.Printf("Response Status: %s\n", resp.Status)
	return resp, nil
}
