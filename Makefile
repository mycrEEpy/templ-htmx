build: generate
	mkdir -p dist
	go build -o ./dist/templ-htmx ./cmd/templ-htmx

generate:
	templ generate ./internal/components

test: generate
	go test -v ./... -race

run: generate
	go run ./cmd/templ-htmx
