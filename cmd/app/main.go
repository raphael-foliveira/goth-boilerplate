package main

import (
	"context"
	"fmt"
	"math/rand"
	"net/http"
	"os"
	"os/signal"

	"github.com/a-h/templ"
	"github.com/labstack/echo/v4"
	"github.com/raphael-foliveira/goth-boilerplate/views"
)

func RenderElement(c echo.Context, status int, component templ.Component) error {
	buf := templ.GetBuffer()
	defer templ.ReleaseBuffer(buf)

	if err := component.Render(c.Request().Context(), buf); err != nil {
		fmt.Println("error rendering component:", err)
		return err
	}

	return c.HTML(status, buf.String())
}

func main() {
	app := echo.New()

	app.Static("/static", "public/assets")

	app.GET("/", func(c echo.Context) error {
		return c.JSON(http.StatusOK, map[string]string{
			"status": "ok",
		})
	})

	app.GET("/hello", func(c echo.Context) error {
		fmt.Println("running hello...")
		component := views.HelloContainer(views.ContainerProps{
			Title: "Foo",
			Name:  "Bar",
		})

		return RenderElement(c, http.StatusOK, component)
	})

	app.POST("/mouse-entered", func(c echo.Context) error {
		fmt.Println("mouse entered")
		randomNumber := rand.Intn(10000)

		return RenderElement(c, http.StatusOK, views.SwappedNumbers(randomNumber))
	})

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
