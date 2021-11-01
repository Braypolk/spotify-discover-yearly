package main

import (
	"fmt"
	"net/url"
	"strings"
)

// get a map of all playlists and associated info
func CheckPlaylists() (map[string]interface{}, error) {
	fmt.Println("retrieving playlists")
	return BuildRequest("GET", "https://api.spotify.com/v1/me/playlists", nil)
}

func CheckSongs(id string) (map[string]interface{}, error) {
	fmt.Println("retrieving songs")
	return BuildRequest("GET", "https://api.spotify.com/v1/playlists/"+id+"/tracks?market=US", nil)
}

// given a playlist name, create an empty playlist
func CreatePlaylist(name string) (string, error) {
	fmt.Println("Creating playlist...")
	body := []byte(`{
		"name": "` + name + `",
		"description": "created with operator"
	  }`)

	response, err := BuildRequest("POST", "https://api.spotify.com/v1/users/ny741pp6gedqst6z0evsk3ymj/playlists", body)
	return response["id"].(string), err
}

func AddSongs(playlist_id string, songs []string) error {
	// TODO: check if songs already exist
	fmt.Println("adding songs...")
	for i := 0; i < len(songs); i++ {
		songs[i] = url.QueryEscape(songs[i])
	}
	result := strings.Join(songs, ",")
	_, err := BuildRequest("POST", "https://api.spotify.com/v1/playlists/"+playlist_id+"/tracks?uris="+result, nil)
	return err
}