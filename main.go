package main

import (
	"net/http"
	"github.com/gin-gonic/gin"
)

func GetRouter() *gin.Engine {
  // making default router
  r := gin.Default()

  // adding static files to the router
  r.Static("/static", "./static")

  // loading all html files
  r.LoadHTMLGlob("templates/*")

  return r
}


func AddMainRoute(title, message string, r *gin.Engine) {
  r.GET("/", func(c *gin.Context) {
    //render the index.html with messages
    c.HTML(http.StatusOK, "index.html", gin.H {
      "title": title,
      "message": message,
    })
  })
}


func AddGetRoute(title, template string, r *gin.Engine) {
  r.GET("/" + title, func(c *gin.Context) {
    c.HTML(http.StatusOK, template, gin.H {
      "title": title,
    })
  })
}


func main() {
  r := GetRouter()

  AddMainRoute("Test", "Hello World", r)

  AddGetRoute("users", "users.html", r)

  r.Run()
}
