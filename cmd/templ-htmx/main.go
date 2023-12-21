package main

import (
	"log/slog"
	"os"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/mycrEEpy/templ-htmx/internal/components"
)

func main() {
	mux := echo.New()

	mux.GET("/", helloHandler)
	mux.GET("/scroll", scrollHandler)

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

func scrollHandler(ctx echo.Context) error {
	page, _ := strconv.Atoi(ctx.QueryParam("page"))

	if page != 0 {
		return components.InfiniteRows(page).Render(ctx.Request().Context(), ctx.Response())
	}

	return components.ScrollPage(page+1).Render(ctx.Request().Context(), ctx.Response())
}
