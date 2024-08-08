package internal

import (
	"encoding/json"
	"io"
	"net/http"
)

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
