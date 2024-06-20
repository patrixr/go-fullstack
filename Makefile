.PHONY: tidy serve build all dev test tailwind clean migrate templ run

# Variables
BIN_NAME := myapp
BUILD_FOLDER := ./.out
VERSION := $(shell cat ./VERSION)

# Default target
all: build

tidy:
	go mod tidy

migrate:
	go run ./ migrate

fmt:
	go fmt ./... && templ fmt .

serve:
	./.out/${BIN_NAME} serve

build: templ tailwind
	go build -o ${BUILD_FOLDER}/${BIN_NAME} ./

run: tailwind
	go run ./ serve

test:
	ENV=test go test -v ./...

templ:
	go run github.com/a-h/templ/cmd/templ@latest generate

clean:
	$(RM) ./**/*_templ.go

tailwind:
	tailwindcss -i public/css/style.css -o public/style.css --minify

dev:
	AIR_ENABLED=1 go run github.com/air-verse/air@latest serve
