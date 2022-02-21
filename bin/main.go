package main

import (
	"fmt"
	_ "net/http"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type Page struct {
	problem     string
	constraints string
}

func main() {
	// DBMigrate(DBconnect())
	r := gin.Default()
	r.LoadHTMLGlob("templates/*")

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	pages := make(map[int]Page)
	pages[0] = Page{"a+bの答えを出力してください", "a,b < 10"}
	pages[1] = Page{"最小の自然数を出力してください", "なし"}

	fmt.Println(pages[0])

	r.GET("/", func(c *gin.Context) {
		c.HTML(200, "index.html", gin.H{})
	})

	r.GET("/problems", func(c *gin.Context) {
		c.HTML(200, "problems.html", gin.H{})
	})

	r.GET("/problems/:i", func(c *gin.Context) {
		i, _ := strconv.Atoi(c.Param("i"))

		c.HTML(200, "problem.html", gin.H{
			"problem": pages[i].problem,
		})
	})

	r.Run(port)
}
