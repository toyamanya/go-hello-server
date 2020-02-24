package add

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo"
)

type jsonData struct {
	Left  int //`json:"left,omitempty"`
	Right int //`json:"right,omitempty"`
}

func main() {
	e := echo.New()
	e.POST("/add", postHandler)
	e.Start(":4000")
}

func postHandler(c echo.Context) error {
	data := new(jsonData)
	err := c.Bind(data)

	// errorHandling
	if err != nil {
		var ans struct {
			Error string `json:"error"`
		}
		ans.Error = "Bad Request"
		return c.JSON(http.StatusBadRequest, ans)
	}

	var ans struct {
		Answer string `json:"answer"`
	}
	ans.Answer = strconv.Itoa(data.Left + data.Right)
	return c.JSON(http.StatusOK, ans)
}

// 構造体の定義でtypeを使う必要はない
/*
	type errorJsonData struct {
		Error string `json:"error"`
	}
	ans := &errorJsonData{
		Error: "Bad Request",
	}

	type answerJsonData struct {
		Answer string `json:"answer"`
	}
	ans := &answerJsonData{
		Answer: tmp,
	}
*/
