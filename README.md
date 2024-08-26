# Alasbot

This is a simple discord bot to let people ask for the status of a 7 days to die server.

## Prerequisited

To run this, you will need:
* A 7 days to die server
* Some place to run the software
* A discord bot token (Needs to be able to read and send messages)

## Running the bot

### Build and run from source

* Compile: `go build -o bot.exe cmd/bot/main.go`
* Set environment variables 
  * `SEVEN_DAYS_SERVER="<hostname>:<port>"`
  * `BOT_TOKEN="Bot <bot token>"`
* Run `bot.exe`

### Running it somewhere else

There is a docker image to run it wherever

```yaml
services:
  alasbot:
    image: ghcr.io/tlanfer/alasbot:main
    restart: always
    environment:
     - BLOODMOON_OFFSET=0
     - BOT_TOKEN=Bot <bot token>
     - SEVEN_DAYS_SERVER=<hostname>:<port>
```


## Adding it to your server

* Use the discord developer portal to generate a link to add the bot to your discord
* Use that link, add the bot to your discord
* It should show as online

## Usage

When you enter `!server` anywhere in your discord it should give you some info.

It should also work in DMs

Any feedback is of course very welcome!
