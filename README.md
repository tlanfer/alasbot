# Alasbot

This is a simple discord bot to let people ask for the status of a 7 days to die server.

## Prerequisited

To run this, you will need:
* A 7 days to die server
* Some place to run the software
* A discord bot token (Needs to be able to read and send messages)

## Running the bot

### Run locally

* Compile: `go build -o bot.exe cmd/bot/main.go`
* Set environment variables 
  * `SEVEN_DAYS_SERVER="<hostname>:<port>"`
  * `BOT_TOKEN="Bot <bot token>"`
* Run `bot.exe`

### Running on heroku

* Check this project out
* Make an account on heroku and install the heroku cli
* Run ``heroku login`` to login (duh)
* Then run

```
# heroku create

Creating app... done, ⬢ mystic-wind-83
Created http://mystic-wind-83.herokuapp.com/ | git@heroku.com:mystic-wind-83.git

# heroku config:set SEVEN_DAYS_SERVER="<hostname>:<port>"
# heroku config:set BOT_TOKEN="Bot <bot token>" 

# git push heroku master

# heroku ps:scale worker=1
```

### Running it somewhere else
If you want to run the bot somewhere else, you're on your own.
It should easily run anywhere you can run linux executable.


## Adding it to your server

* Use the discord developer portal to generate a link to add the bot to your discord
* Use that link, add the bot to your discord
* It should show as online

## Usage

When you enter `!server` anywhere in your discord it should give you some info.

It should also work in DMs

Any feedback is of course very welcome!