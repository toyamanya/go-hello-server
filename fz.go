package fz

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo"
)

func fz() {
	e := echo.New()

	e.GET("/fizzbuzz", fizzbuzz)
	e.Start(":4000")
}

func fizzbuzz(c echo.Context) error {
	count, err := strconv.Atoi(c.QueryParam("count"))

	if err != nil {
		return c.String(http.StatusBadRequest, "count is not correct")
	}
	if count <= 0 {
		return c.String(http.StatusBadRequest, "count is positive")
	}
	return c.String(http.StatusOK, calcFizzbuzz(count))
}

func calcFizzbuzz(num int) string {
	i := 1
	str := ""
	for i < num+1 {
		switch {
		case i%15 == 0:
			str += "FIZZ BUZZ\n"
		case i%3 == 0:
			str += "FIZZ\n"
		case i%5 == 0:
			str += "BUZZ\n"
		default:
			str += strconv.Itoa(i) + "\n"
		}
		i++
	}
	return str
}
