# Spotify Discover Yearly

:exclamation:**Please note the Go version is the more complete version of this project and should be used as the default. The JS version should be used for fun and not an actual deployment**:exclamation:

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
    - to build for different runtimes (helpful for deploying to cloud) use `env GOOS=linux GOARCH=arm64 go build .` where linux and arm64 would be the params you change out
    - you can check current params with `go env GOOS GOARCH`
    - check list of different configs with `go tool dist list`
- Then you will have an executable, so you can run `./spotify-discover-yearly` on Mac/Linux or `./spotify-discover-yearly.exe` on Windows

This should authenticate with previously retreieved refresh token that you definitely already put in the env files. Then it will extract songs from your discover weekly playlist and put them into your fancy new discover yearly playlist.

Unless you want to remember to manually run this once a week it would be best to put on a cron job on a raspberry pi or somewhere in the cloud.

## How I automated it
I used [AWS](https://aws.amazon.com/), specifically their lambda functions
- Create a new function
    - author from scratch
    - python runtime
    - I used arm64 architecture because why not (this is where you would make sure you have the correct runtime if you use go)

I followed [this article](https://medium.com/@biancanhinojosa/running-executables-in-aws-lambda-dc79b8f33ec7). Which is basically just `aws/labmba_function.py` and the commands in `awszip.sh`

**So if you finished the Setup and Next Steps portions and created a labmda function, you should just be able to run `awszip.sh` and upload the zip to your lambda function.**

if you cannot run awszip.sh, run `chmod 755 awszip.sh`

I did have some weird issues with google chrome when trying to upload the zip file. So if it's not letting you, try a different browser. Based on AWS [pricing docs](https://aws.amazon.com/lambda/pricing/) this is well under the free tier limits. So now we have easy automation for free! (or at least hopefully, you should check the docs to make sure pricing hasn't changed)


---

Todo:
- check if yearly playlist already exists, if not create new one
- auto populate new refresh token to env file from running get-new-refresh.js
- check for duplicate songs before adding
- it would be nice to just have one env file
- notification when finished (prob just an email)
