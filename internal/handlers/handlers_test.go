package handlers_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/raphael-foliveira/goth-boilerplate/internal/handlers"
	"github.com/stretchr/testify/assert"
)

func setUp(t *testing.T) *echo.Echo {
	app := echo.New()
	handlers.MountRoutes(app)
	t.Cleanup(func() {})
	return app
}

func TestHandlers_Hello(t *testing.T) {
	app := setUp(t)

	req := httptest.NewRequest(http.MethodGet, "/hello", nil)
	recorder := httptest.NewRecorder()

	app.ServeHTTP(recorder, req)

	assert.Contains(t, recorder.Body.String(), "Hello, Bar!")
}

func TestHandlers_Healthcheck(t *testing.T) {
	app := setUp(t)

	req := httptest.NewRequest(http.MethodGet, "/", nil)
	recorder := httptest.NewRecorder()

	app.ServeHTTP(recorder, req)

	assert.Equal(t, http.StatusOK, recorder.Code)
}
