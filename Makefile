NAME=dodo-apex-robot
RELEASE_DIR=release
GOBUILD=CGO_ENABLED=0 go build -trimpath -ldflags '-w -s'

all: release
.PHONY: release

release:
	GOARCH=amd64 GOOS=linux $(GOBUILD) -o $(RELEASE_DIR)/$(NAME)
	cp example.config.json $(RELEASE_DIR)/config.json

clean:
	rm $(RELEASE_DIR)/*