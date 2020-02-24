package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"

	"github.com/labstack/echo"
)

type Student struct {
	Number int    `json:"student_number"`
	Name   string `json:"name"`
}

type Class struct {
	Number   int       `json:"class_number"`
	Students []Student `json:"students"`
}

func main() {
	e := echo.New()

	e.GET("/students/:class/:studentNumber", dataHandler)
	e.Start(":4000")
}

func dataHandler(c echo.Context) error {
	c_num, err := strconv.Atoi(c.Param("class"))
	s_num, err := strconv.Atoi(c.Param("studentNumber"))

	// 文字が入力されたらエラーを返す
	if err != nil {
		return c.String(http.StatusBadRequest, "class_number and student_num are number ")
	}

	// JSONファイルの読み込み
	bytes, err := ioutil.ReadFile("data.json")
	if err != nil {
		log.Fatal(err)
	}

	// JSONデコード
	var classes []Class
	if err := json.Unmarshal(bytes, &classes); err != nil {
		log.Fatal(err)
	}
	// 今JSONのデータを構造体として持っている

	// デコードしたデータに当てはまるものがあればStudent structを返す
	for _, val := range classes {
		// c_numに一致するものがある
		if val.Number == c_num {
			//s_numが一致
			for _, elm := range val.Students {
				if elm.Number == s_num {
					// 見つかった場合
					return c.JSONPretty(http.StatusOK, elm, "    ")
				}
			}
		}
	}

	// 見つからなかった場合
	var ans struct {
		Error string `json:"error"`
	}
	ans.Error = "Student Not Found"
	return c.JSONPretty(http.StatusBadRequest, ans, "    ")

}
