package main

import (
	"fmt"
	"log"
	"strconv"
	"time"
)

var AccessToken string

var BasicCreds string = ClientId + ":" + ClientSecret
var ApiUrl string = "https://api.spotify.com/v1/"
var UserId string = ""

func main() {
	var err error
	UserId, err = GetUser()
	if err != nil {
		log.Fatal(err)
	}

	// get all playlists in account
	playlists, err := CheckPlaylists()
	if err != nil {
		log.Fatal(err)
	}

	var disco_week_id string
	var disco_year_id string

	// search through playlists to find Discover Weekly and this year's Discover Yearly
	for _, p := range playlists["items"].([]interface{}) {
		name := p.(map[string]interface{})["name"].(string)
		id := p.(map[string]interface{})["id"].(string)
		if name == "Discover Weekly" {
			disco_week_id = id
		} else if name == (strconv.Itoa(time.Now().Year()) + " Discover Yearly") {
			disco_year_id = id
		}
	}

	// in case discover weekly doesn't exist for some reason
	if disco_week_id == "" {
		log.Fatal("Could not find discover weekly playlist, be sure it is added to your library by hearting the playlist")
	}
	// create yearly playlist if it doesn't exist
	if disco_year_id == "" {
		disco_year_id, err = CreatePlaylist(strconv.Itoa(time.Now().Year()) + " Discover Yearly")
		if err != nil {
			log.Print("ERROR ON CREATEPLAYLIST: ")
			log.Fatal(err)
		}
		fmt.Println("Created new playlist. Playlist ID: " + disco_year_id)
	}

	songs_to_add := FormattedPlaylistSongs(disco_week_id)
	songs_in_yearly := FormattedPlaylistSongs(disco_year_id)

	// remove any duplicates in songs_to_add
	for id, _ := range songs_to_add {
		if _, ok := songs_in_yearly[id]; ok {
			delete(songs_to_add, id)
		}
	}

	// finally, add songs to discover yearly playlist
	if len(songs_to_add) != 0 {
		err := AddSongs(disco_year_id, songs_to_add)
		if err != nil {
			log.Fatal(err)
		}

		// WriteSliceToFile(song_names, "pastSongs/"+strconv.Itoa(time.Now().Year())+".yaml")
		for _, i := range songs_to_add {
			fmt.Println(i)
		}
	} else {
		fmt.Println("No songs to add")
	}
}
