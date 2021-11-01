var request = require('request');
var a = require("./env.js");

var y = (new Date().getFullYear()).toString()
var client_id = a.client_id;
var client_secret = a.client_secret;
var refresh_token = a.refresh_token;

var authOptions = {
    url: 'https://accounts.spotify.com/api/token',
    headers: { 'Authorization': 'Basic ' + (new Buffer(client_id + ':' + client_secret).toString('base64')) },
    form: {
        grant_type: 'refresh_token',
        refresh_token: refresh_token
    },
    json: true
};

request.post(authOptions, function(error, response, body) {
    if (!error && response.statusCode === 200) {

        var access_token = body.access_token;

        var getPlaylists = {
            url: 'https://api.spotify.com/v1/me/playlists',
            headers: { 'Authorization': 'Bearer ' + access_token },
            json: true
        };

        var playlists = [];
        var discoWeek = '';
        var discoYear = '';
        request.get(getPlaylists, function(error, response, body) {
            if (error || response.statusCode != 200) {
                console.log("GETPLAYLISTS BAD");
                console.log(response.statusCode);
                console.log(body);
            }

            playlists = body.items;
            for (const e of playlists) {
                if (e.name === 'Discover Weekly') {
                    discoWeek = e.id;
                } else if (e.name.indexOf('Discover Yearly') !== -1) {
                    if (e.name.indexOf(year) !== -1) {
                        discoYear = e.id;
                    } else {

                    }
                }
            }

            var getTracks = {
                url: `https://api.spotify.com/v1/playlists/${discoWeek}/tracks?market=US`,
                headers: { 'Authorization': 'Bearer ' + access_token },
                market: 'US',
                json: true
            }

            request.get(getTracks, function(error, response, body) {
                if (error || response.statusCode != 200) {
                    console.log("GETTRACKS BAD");
                    console.log(response.statusCode);
                    console.log(body);
                }

                var tracks = body.items;
                var trackUris = '';

                for (const e of tracks) {
                    trackUris += e.track.uri + ',';
                }
                trackUris = trackUris.substring(0, trackUris.length - 1);

                var headers = {
                    'Accept': 'application/json',
                    'Content-Type': 'application/json',
                    'Authorization': `Bearer ${access_token}`
                };

                var options = {
                    url: `https://api.spotify.com/v1/playlists/${discoYear}/tracks?uris=${trackUris}`,
                    method: 'POST',
                    headers: headers,
                };

                request(options, function(error, response, body) {
                    if (error || (response.statusCode != 200 && response.statusCode != 201)) {
                        console.log("ADD SONGS BAD");
                        console.log(response.statusCode);
                        console.log(body);
                    } else {
                        console.log("you done did it. Enjoy them new songs");
                    }
                })
            });
        });
    }
});
console.log("DONE");