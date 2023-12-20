package main

import (
	"log/slog"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/mycrEEpy/templ-htmx/internal/components"
)

func main() {
	mux := echo.New()

	mux.GET("/", helloHandler)

	mux.HideBanner = true
	mux.HidePort = true

	err := mux.Start(":8000")
	if err != nil {
		slog.Error("mux failure", slog.String("err", err.Error()))
		os.Exit(1)
	}
}

func helloHandler(ctx echo.Context) error {
	return components.Hello("world").Render(ctx.Request().Context(), ctx.Response())
}
