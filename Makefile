BINARY_NAME		= twitch-bot
BINARY_PATH		= bin

default: build

test:
	echo "No tests right now"
	
build:
	mkdir -p $(CURDIR)/$(BINARY_PATH)
	go build -o $(CURDIR)/$(MODULE_BINARY_PATH)/$(BINARY_NAME) main.go