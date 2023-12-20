build: generate
	mkdir -p dist
	go build -o ./dist/templ-htmx ./cmd/templ-htmx

generate:
	templ generate ./internal/components

test:
	go test -v ./... -race

run:
	go run ./cmd/templ-htmx
