package api

import (
  "net/http"

  "github.com/gin-gonic/gin"

  "fmt"
)


type User struct {
  Name string
  Email string
}

func HandlePostUsers(r *gin.Engine, users []User) {
  r.POST("/add-user", func(c *gin.Context) {
    name := c.PostForm("name")
    email := c.PostForm("email")

    newUser := User {
      Name: name,
      Email: email,
    }

    users = append(users, newUser)

    fmt.Println(users)
    c.String(http.StatusOK, "User %s succesfully added", newUser.Name)
  })
}
