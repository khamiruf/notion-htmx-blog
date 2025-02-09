.PHONY: dev build css watch-css install

install:
	go mod tidy
	npm install

dev: watch-css
	go run cmd/server/main.go

build:
	npm run build
	go build -o bin/server cmd/server/main.go

css:
	npm run build

watch-css:
	npm run dev & 

clean:
	rm -rf bin/
	rm -f web/static/css/output.css