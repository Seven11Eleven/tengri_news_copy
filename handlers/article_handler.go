package handlers

import (
    "net/http"

    "github.com/gin-gonic/gin"
    "tengri_news/parser"
)


 

    func SetupRouter() *gin.Engine {
	r := gin.Default()

    // r.Static("/static", "./static")
	r.LoadHTMLGlob("../templates/*.html")

    r.GET("/", func(c *gin.Context) {
        // Возвращаем HTML-шаблон для главной страницы новостей
        c.HTML(http.StatusOK, "news.html", nil)
    })


    r.GET("/news", parsernews.ParseLastNews)
    r.GET("/news/*url", parsernews.ParseOneNewsByUrl)


	return r
}