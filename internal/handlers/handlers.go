package handlers

import (
	"fmt"
	"math/rand"
	"net/http"

	"github.com/a-h/templ"
	"github.com/labstack/echo/v4"
	"github.com/raphael-foliveira/goth-boilerplate/views"
	"github.com/raphael-foliveira/goth-boilerplate/views/components"
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

func Hello(c echo.Context) error {
	fmt.Println("running hello...")
	component := views.HelloContainer(views.ContainerProps{
		Title: "Foo",
		Name:  "Bar",
	})

	return RenderElement(c, http.StatusOK, component)
}

func Click(c echo.Context) error {
	randomNumber := rand.Intn(10000)

	return RenderElement(c, http.StatusOK, components.SwappedNumbers(randomNumber))
}

func Healthcheck(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]string{
		"status": "ok",
	})
}

func MountRoutes(app *echo.Echo) {
	app.Static("/static", "public/assets")
	app.GET("/", Healthcheck)
	app.GET("/hello", Hello)
	app.POST("/mouse-clicked", Click)
}
