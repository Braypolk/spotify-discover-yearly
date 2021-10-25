# Spotify Discover Yearly

### Setup
If refresh token has not been created, has somehow expired, or just isn't working, run `node get-new-refresh.js`
Make sure other vars in `env.js` are set to the correct values

### Main Stuff
Run main file: `node app.js`
This will authenticate with previously retreieved refresh token located in `env.js`. Then it will extract songs from your discover weekly playlist and put them into your discovery yearly playlist.

Unless you want to remember to run this once a week it would be best to put on a cron job on a raspberry pi or cloud.

---

Todo:
- check if yearly playlist already exists, if not create new one
- auto populate new refresh token to env file from running get-new-refresh.js
- check for duplicate songs before adding
