package main

import (
	"net/http"

	"github.com/gin-gonic/gin"

  "main/api"
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


func AddUsersRoute(r *gin.Engine, users *[]api.User) {
  r.GET("/users", func(c *gin.Context) {
    c.HTML(http.StatusOK, "users.html", gin.H {
      "Users": *users,
      "title": "users",
    })
  })
} 

func AddUserPost (r *gin.Engine, users *[]api.User) {
  r.POST("/add-user", func(c *gin.Context) {
    name := c.PostForm("name")
    email := c.PostForm("email")

    newUser := api.User {
      Name: name,
      Email: email,
    }

    *users = append(*users, newUser)

    //c.HTML(http.StatusOK, "users.html", gin.H {
    //  "Users": *users,
    //  "title": "users",
    //})
    c.Redirect(http.StatusFound, "/users")
  })
}

func main() {
  r := GetRouter()
  users := []api.User{
    {"Bob", "bob@gmail.com"},
    {"Alice", "aliceIsTesting@mail.com"},
    {"Karl", "karlLovesPizza@nyy.com"},
  } 

  AddMainRoute("Test", "Hello World", r)

  AddGetRoute("add-users", "add-users.html", r)

  AddUsersRoute(r, &users)
  
  AddUserPost(r, &users)

  r.Run()
}
