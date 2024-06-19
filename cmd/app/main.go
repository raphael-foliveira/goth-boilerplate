package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"

	"github.com/labstack/echo/v4"
	"github.com/raphael-foliveira/goth-boilerplate/internal/handlers"
)

func startServer(app *echo.Echo) {
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()

	go func() {
		if err := app.Start(":3000"); err != nil {
			fmt.Println(err)
		}
	}()

	<-ctx.Done()
	fmt.Println("\nServer shutting down")
}

func main() {
	app := echo.New()
	handlers.MountRoutes(app)
	startServer(app)
}
