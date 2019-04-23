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

# Docker
Build
```bash
docker build -t twitch-bot .
```
List docker image
```bash
docker image ls
```
Run
```bash
docker run --name twitch-bot twitch-bot
```
Stop
```bash
docker stop twitch-bot
```
Remove container
```bash
docker rm twitch-bot
```
Remove all inactive or stopped containers
```bash
docker rm $(docker ps -aq)
```
