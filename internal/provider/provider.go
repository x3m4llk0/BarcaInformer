package provider

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"
)

type Result struct {
	Filters   map[string]interface{} `json:"filters"`
	ResultSet map[string]interface{} `json:"resultSet"`
	Matches   []Matches              `json:"matches"`
}

type Matches struct {
	Area        map[string]interface{}   `json:"area"`
	Competition Competition              `json:"competition"`
	Season      map[string]interface{}   `json:"season"`
	Id          int                      `json:"id"`
	UtcDate     time.Time                `json:"utcDate"`
	Status      string                   `json:"status"`
	MatchDay    int                      `json:"matchday"`
	Stage       string                   `json:"stage"`
	Group       string                   `json:"group"`
	LastUpdated time.Time                `json:"lastUpdated"`
	HomeTeam    Team                     `json:"homeTeam"`
	AwayTeam    Team                     `json:"awayTeam"`
	Score       map[string]interface{}   `json:"score"`
	Odds        map[string]interface{}   `json:"odds"`
	Referees    []map[string]interface{} `json:"referees"`
}

type Team struct {
	Id        int    `json:"id,omitempty"`
	Name      string `json:"name,omitempty"`
	ShortName string `json:"shortName,omitempty"`
	Tla       string `json:"tla,omitempty"`
	Crest     string `json:"crest" json:"crest,omitempty"`
}

type Competition struct {
	Id     int    `json:"id"`
	Name   string `json:"name"`
	Code   string `json:"code"`
	Type   string `json:"type"`
	Emblem string `json:"emblem"`
}

type GetInfoResult struct {
	Competition Competition `json:"competition"`
	HomeTeam    Team        `json:"homeTeam"`
	AwayTeam    Team        `json:"awayTeam"`
	UtcDate     time.Time   `json:"utcDate"`
}

func GetInfo() (GetInfoResult, error) {
	url := "https://api.football-data.org/v4/teams/81/matches?status=SCHEDULED&limit=1"
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

	body, err := io.ReadAll(resp.Body)

	var unmarshresult Result
	json.Unmarshal(body, &unmarshresult)

	var GetInfoResult GetInfoResult
	json.Unmarshal(body, &GetInfoResult)

	GetInfoResult.Competition = unmarshresult.Matches[0].Competition
	GetInfoResult.HomeTeam = unmarshresult.Matches[0].HomeTeam
	GetInfoResult.AwayTeam = unmarshresult.Matches[0].AwayTeam
	GetInfoResult.UtcDate = unmarshresult.Matches[0].UtcDate

	// Handle the response
	fmt.Printf("Response Status: %s\n", resp.Status)
	return GetInfoResult, nil
}
