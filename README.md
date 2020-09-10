# Alasbot

This is a simple discord bot to let people ask for the status of a 7 days to die server.

## Prerequisited

To run this, you will need:
* A 7 days to die server
* Some place to run the software
* A discord bot token (Needs to be able to read and send messages)

## Run locally

* Compile: `go build -o bot.exe cmd/bot/main.go`
* Set environment variables 
  * `SEVEN_DAYS_SERVER="<hostname>:<port>"`
  * `BOT_TOKEN="Bot <bot token>"`
* Run `bot.exe`

## Running on heroku

* Check this project out
* Make an account on heroku and install the heroku cli
* Run ``heroku login`` to login (duh)
* Then run

```
heroku create
Creating app... done, ⬢ mystic-wind-83
Created http://mystic-wind-83.herokuapp.com/ | git@heroku.com:mystic-wind-83.git

heroku config:set SEVEN_DAYS_SERVER="<hostname>:<port>"
heroku config:set BOT_TOKEN="Bot <bot token>" 

git push heroku master

heroku ps:scale worker=1
```

* Add the bot to your server. It should show as online

Now, when you enter !server it should give you some info.
It should also work in DMs