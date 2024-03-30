run: build
	./bin/app

install:
	go install github.com/a-h/templ/cmd/templ@latest
	go get ./...
	go mod vendor
	go mod tidy
	go mod download
	npm install -D tailwindcss
	npm install -D daisyui@latest

build:
	templ generate view
	go build -tags dev -o bin/app main.go