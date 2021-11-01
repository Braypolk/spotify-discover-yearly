package main
import (
	"fmt"
	"log"
	"strings"
	"time"
	"strconv"
)

var access_token string

var BasicCreds string = ClientId +":"+ ClientSecret
var ApiUrl string = "https://accounts.spotify.com/api/token"
	
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
		if(strings.Contains(name, "Discovery Weekly")){
			disco_week_id = id
		} else if (strings.Contains(name, "Discover Yearly")) {
			if (strings.Contains(name, strconv.Itoa(time.Now().Year()))) {
				disco_year_id = id;
			}
		}
	}

	fmt.Println(disco_year_id)

	// TODO: if discoyear or week nil do something about it

	songs, err := CheckSongs(disco_week_id)
	var songs_to_add []string


	// store each song from playlist in spotify to songs_to_add
	for _, item := range songs["items"].([]interface{}) {
		// fmt.Println(item.(map[string]interface{}))
		// name := item.(map[string]interface{})["name"].(string)
		// id := item.(map[string]interface{})["id"].(string)
		uri := item.(map[string]interface{})["uri"].(string)
		songs_to_add = append(songs_to_add, uri)
	}

	if len(songs_to_add) != 0 {
		err := AddSongs("1Cq8eVWuqx0RJ4XHMUlghj", songs_to_add)
		if err != nil {
			log.Fatal(err)
		}
	}
}
