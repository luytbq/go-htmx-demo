run: build
	./bin/app

air: 
	@air

templ:
	@templ generate --watch --proxy=http://localhost:3000

css:
	@tailwindcss -i view/css/app.css -o static/styles.css --watch 

build:
	@tailwindcss -i view/css/app.css -o static/styles.css
	templ generate view
	go build -tags dev -o bin/app main.go

install:
	go install github.com/a-h/templ/cmd/templ@latest
	go get ./...
	go mod vendor
	go mod tidy
	go mod download
	npm install -D tailwindcss postcss autoprefixer
	npm install -D daisyui@latest
