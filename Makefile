build: tidy generate
	mkdir -p dist
	go build -o ./dist/templ-htmx ./cmd/templ-htmx

generate:
	templ generate ./internal/components

test: generate
	go test -v ./... -race

run: tidy generate
	go run ./cmd/templ-htmx

tidy:
	go mod tidy
