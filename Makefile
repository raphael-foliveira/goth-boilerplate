dev:
	npm run css | air

build:
	go build -o ./bin/main ./cmd/app

run: build
	./bin/main
