# twitch-bot
A twitch.tv bot where users can play text games and grab statistics.

# Getting Started
```bash
$ git clone github.com/hueyjj/twitch-bot
```
or
```bash
$ go get -u github.com/hueyjj/twitch-bot
```

# Setup .env
```bash
BOT_USERNAME=your_username
CHANNEL_NAME=channel_name_to_join
OAUTH_TOKEN=generate_token_from_twitch_api
```

# Running
```bash
$ cd github.com/hueyjj/twitch-bot
$ make
$ ./bin/twitch-bot
```


# Cleaning
```bash
$ make clean
```