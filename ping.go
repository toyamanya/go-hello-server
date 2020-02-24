package ping

import (
	"net/http"

	"github.com/labstack/echo"
)

type jsonData struct {
	Number int    `json:"number, omitempty"`
	String string `json:"string,omitempty"`
	Bool   bool   `json:bool,omitempty"`
}

func main() {
	e := echo.New()

	e.GET("/ping", func(c echo.Context) error {
		return c.String(http.StatusOK, "pong")
	})
	e.Start(":4000")
}
