package main

import (
	"fmt"
	"log"
	// "reflect"
	"strconv"
	"strings"
	"time"
)

var AccessToken string

var BasicCreds string = ClientId +":"+ ClientSecret
var ApiUrl string = "https://api.spotify.com/v1/"
	
func main() {
	fmt.Println(BasicCreds)
	fmt.Println(RefreshToken)
	playlists, err := CheckPlaylists()
	if err != nil {
		log.Fatal(err)
	}

	var disco_week_id string
	var disco_year_id string

	// convert to slice for subset check
	for _, p := range playlists["items"].([]interface{}) {
		name := p.(map[string]interface{})["name"].(string)
		id := p.(map[string]interface{})["id"].(string)
		if(strings.Contains(name, "Discover Weekly")){
			disco_week_id = id
		} else if (strings.Contains(name, "Discover Yearly")) {
			if (strings.Contains(name, strconv.Itoa(time.Now().Year()))) {
				disco_year_id = id;
			}
		}
	}

	// TODO: if discoyear or week nil do something about it

	songs, err := CheckSongs(disco_week_id)
	if err != nil {
		log.Print("ERROR ON CHECKSONGS: ")
		log.Fatal(err)
	}

	var songs_to_add []string
	var song_names []string

	// store each song from playlist in spotify to songs_to_add
	for _, item := range songs["items"].([]interface{}) {
		name := (item.(map[string]interface{})["track"].(map[string]interface{})["name"]).(string)
		uri := (item.(map[string]interface{})["track"].(map[string]interface{})["uri"]).(string)
		song_names = append(song_names, name)
		songs_to_add = append(songs_to_add, uri)
	}

	// keep text log of file names
	WriteSliceToFile(song_names, strconv.Itoa(time.Now().Year())+".yaml")

	// finally, add songs to discover yearly playlist
	if len(songs_to_add) != 0 {
		err := AddSongs(disco_year_id, songs_to_add)
		if err != nil {
			log.Fatal(err)
		}
	}
}
