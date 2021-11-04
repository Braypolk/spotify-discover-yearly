# Spotify Discover Yearly

## Spotify Setup 
1. Sign up for / Log in to [Spotify Dev Account](https://developer.spotify.com/dashboard) with the Spotify account you want to use
2. Create a new app
3. Click 'Edit Settings' button and add http://localhost:8888/callback/ to 'Redirect URIs'
    - if you end up using a different redirect uri be sure to change `redirect_uri` in get-new-refresh.js
4. Copy Client ID and Client Secret to env files


## Next Steps
If refresh token has not been created, somehow expired, or just isn't working, run `node js/get-new-refresh.js` and go to http://localhost:8888

This will go through the authentication process with your account so the api has authorization to interact with your account.
You will then see your refresh token, copy that to env files. 

**DON'T SHARE THIS REFRESH TOKEN OR PUSH IT TO REPO** anyone who has this key will be able to view and edit pretty much anything in your spotify account until a new refresh key is generated.

Note: refresh token should last forever, so you should only need to do this once. But sometimes things act weird so this is an easy way to get a new one

I think that's all the setup required


## Main Stuff
For JS solution, run `node app.js`
For Go solution, 
- For the first time or if you made any changes to the go files, run `go build .`
- Then you will have an executable, so you can do `./spotify-discover-yearly` on Mac/Linux or `./spotify-discover-yearly.exe` on Windows

This should authenticate with previously retreieved refresh token that you definitely already put in the env files. Then it will extract songs from your discover weekly playlist and put them into your fancy new discover yearly playlist.

Unless you want to remember to manually run this once a week it would be best to put on a cron job on a raspberry pi or somewhere in the cloud.

---

Todo:
- check if yearly playlist already exists, if not create new one
- auto populate new refresh token to env file from running get-new-refresh.js
- check for duplicate songs before adding
- it would be nice to just have one env file
- notification when finished (prob just an email)