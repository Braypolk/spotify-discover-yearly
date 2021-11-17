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

func AddSongs(playlist_id string, songs []string) error {
	// TODO: check if songs already exist
	// fmt.Println("adding songs...")
	for i := 0; i < len(songs); i++ {
		songs[i] = url.QueryEscape(songs[i])
	}
	result := strings.Join(songs, ",")
	_, err := BuildRequest("POST", ApiUrl+"playlists/"+playlist_id+"/tracks?uris="+result, nil)
	return err
}
