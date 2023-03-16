package cmd

import (
	"encoding/json"
	"io"
	"net/http"
)

type Artist struct {
	ID             int      `json:"id"`
	Image          string   `json:"image"`
	Name           string   `json:"name"`
	Members        []string `json:"members"`
	CreationDate   int      `json:"creationDate"`
	FirstAlbum     string   `json:"firstAlbum"`
	Locations      string   `json:"locations"`
	ConcertDates   string   `json:"concertDates"`
	Relations      string   `json:"relations"`
	DatesLocations map[string][]string
}

var (
	artists   = []Artist{}
	relations = Relation{}
)

type Relation struct {
	Index []struct {
		ID             int `json:"id"`
		DatesLocations map[string][]string
	}
}

func GetData() error {
	if len(artists) != 0 {
		return nil
	}
	resp, err := http.Get("https://groupietrackers.herokuapp.com/api/artists")
	if err != nil {
		return err
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	err = json.Unmarshal(body, &artists)
	if err != nil {
		return err
	}
	return nil
}

func GetRel() error {
	resp, err := http.Get("https://groupietrackers.herokuapp.com/api/relation")
	if err != nil {
		return err
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	err = json.Unmarshal(body, &relations)
	if err != nil {
		return err
	}
	for i := range artists {
		artists[i].DatesLocations = relations.Index[i].DatesLocations
	}
	return nil
}
