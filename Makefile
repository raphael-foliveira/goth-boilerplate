test:
	go test -v ./...

templ:
	templ generate

test-cover:
	go test ./... -v -coverprofile=c.out -coverpkg=./internal/handlers
	go tool cover -html=c.out
	rm c.out

test-watch:
	air -c .air.test.toml

dev:
	air -c .air.toml

build: templ
	go build -o ./bin/main ./cmd/app

run: build
	./bin/main
