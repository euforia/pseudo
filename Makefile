
NAME = pseudo
COMMIT = $(shell git rev-parse --short HEAD 2> /dev/null || date '+%s')
VERSION = $(shell git describe 2> /dev/null || echo "0.0.0-$(COMMIT)")
BUILDTIME = $(shell date +%Y-%m-%dT%T%z)


DIST_OPTS = -a -tags netgo -installsuffix netgo
LD_OPTS = -ldflags="-X main.version=$(VERSION) -X main.buildtime=$(BUILDTIME) -w"

BUILD_CMD = CGO_ENABLED=0 go build $(LD_OPTS)

SOURCE_FILES = ./cmd/*.go

clean:
	rm -f ./$(NAME)
	rm -rf dist

deps:
	go get github.com/golang/dep/cmd/dep
	dep ensure

test:
	go test -cover $(shell go list ./... | grep -v /vendor | grep -v /cmd)

pseudo:
	$(BUILD_CMD) -o $(NAME) $(SOURCE_FILES)

dist:
	mkdir dist
	GOOS=linux $(BUILD_CMD) $(DIST_OPTS) -o dist/$(NAME)-linux $(SOURCE_FILES)
	GOOS=darwin $(BUILD_CMD) $(DIST_OPTS) -o dist/$(NAME)-darwin $(SOURCE_FILES)
	cd dist; tar -czf $(NAME)-$(VERSION)-linux.tgz $(NAME)-linux; cd -
	cd dist; tar -czf $(NAME)-$(VERSION)-darwin.tgz $(NAME)-darwin; cd -
