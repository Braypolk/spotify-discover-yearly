package main

import (
	// "fmt"
	"net/url"
	"strings"
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

func CheckSongs(id string) (map[string]interface{}, error) {
	// fmt.Println("retrieving songs")
	return BuildRequest("GET", ApiUrl+"playlists/"+id+"/tracks?market=US", nil)
}

// given a playlist name, create an empty playlist
func CreatePlaylist(name string) (string, error) {
	// fmt.Println("Creating playlist...")
	body := []byte(`{
		"name": "` + name + `",
		"description": "compilaion of your discovor weekly songs"
	  }`)

	//   BUG: need to use actual user instead of hardcode
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
