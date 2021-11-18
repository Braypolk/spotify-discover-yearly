package main

import (
	// "fmt"
	"net/url"
	"strconv"
	"strings"
	"math"
)

func GetUser() (string, error) {
	response, err := BuildRequest("GET", ApiUrl+"me", nil)
	return response["id"].(string), err
}

// get a map of all playlists and associated info
func CheckPlaylists() (map[string]interface{}, error) {
	// fmt.Println("retrieving playlists")
	return BuildRequest("GET", ApiUrl+"me/playlists", nil)
}

// Get all items in a playlist
func CheckSongs(id string) (map[string]interface{}, error) {
	// fmt.Println("retrieving songs")
	mp, err := BuildRequest("GET", ApiUrl+"playlists/"+id+"/tracks?market=US", nil)
	if err != nil {
		return mp, err
	}

	limit := mp["limit"].(float64)
	pages := int(math.Ceil((mp["total"].(float64) - float64(len(mp["items"].([]interface{})))) / limit))

	// if the playlist has more items than the api response limit, make multiple calls until all items have been recieved
	for offset := 1; offset <= pages; offset++ {
		url := ApiUrl+"playlists/"+id+"/tracks?market=US&offset="+strconv.Itoa((int(limit)*offset))

		response, err := BuildRequest("GET", url, nil)
		if err != nil {
			return response, err
		}

		mp["items"] = append(mp["items"].([]interface{}), response["items"].([]interface{})...)
	}
	return mp, err
}

// given a playlist name, create an empty playlist
func CreatePlaylist(name string) (string, error) {
	// fmt.Println("Creating playlist...")
	body := []byte(`{
		"name": "` + name + `",
		"description": "compilaion of your discovor weekly songs"
	  }`)
	  
	  response, err := BuildRequest("POST", ApiUrl+"users/"+UserId+"/playlists", body)
	return response["id"].(string), err
}

func AddSongs(playlist_id string, songs map[string]string) error {
	// fmt.Println("adding songs...")
	var ids []string

	for id, _ := range songs {
		songs[id] = url.QueryEscape(songs[id])
		ids = append(ids, id)
	}

	result := strings.Join(ids, ",")
	_, err := BuildRequest("POST", ApiUrl+"playlists/"+playlist_id+"/tracks?uris="+result, nil)
	return err
}
