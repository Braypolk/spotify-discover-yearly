package main

import "log"

// get songs 
func FormattedPlaylistSongs(playlist_id string) map[string]string {
	songs, err := CheckSongs(playlist_id)
	if err != nil {
		log.Print("ERROR ON CHECKSONGS: ")
		log.Fatal(err)
	}

	songs_map := make(map[string]string)

	// store each song id and name from playlist in spotify to songs_to_add
	for _, item := range songs["items"].([]interface{}) {
		name := (item.(map[string]interface{})["track"].(map[string]interface{})["name"]).(string)
		uri := (item.(map[string]interface{})["track"].(map[string]interface{})["uri"]).(string)
		songs_map[uri] = name
	}

	return songs_map
}