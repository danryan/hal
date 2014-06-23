PKG_NAME := hal

all: deps build

deps:
	@go get -d -v ./...

build:
	@mkdir -p bin
	@go build -o bin/$(PKG_NAME)

install:
	@go install $(PKG_NAME)

clean:
	@go clean $(PKG_NAME)
	
.PHONY: all deps build
