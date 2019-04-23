BINARY_NAME		= twitch-bot
BINARY_PATH		= bin

default: build

test:
	echo "No tests right now"
	
build:
	mkdir -p $(CURDIR)/$(BINARY_PATH)
	GO111MODULE=on go build -o $(CURDIR)/$(BINARY_PATH)/$(BINARY_NAME) main.go

clean:
	rm $(CURDIR)/$(BINARY_PATH)/$(BINARY_NAME)