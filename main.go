package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"
)

// User is the representation of a client.
type User struct {
	UUID      string `json:"uuid"`
	FirstName string `json:"first_name"`
	LastName  string `json:"-"`
}

var listUser = map[string]User{}

func main() {
	r := gin.Default()
	r.GET("/users/:uuid", GetUserHandler)
	r.GET("/users", GetAllUserHandler)
	r.POST("/users", PostUserHandler)
	r.Run(":8080")
}

// GetUserHandler is retriving user from the given uuid param.
func GetUserHandler(ctx *gin.Context) {
	if u, ok := listUser[ctx.Param("uuid")]; ok {
		ctx.JSON(http.StatusOK, u)
		return
	}
	ctx.JSON(http.StatusNotFound, nil)
}

// GetAllUserHandler is retriving all users from the database.
func GetAllUserHandler(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, listUser)
}

// PostUserHandler is creating a new user into the database.
func PostUserHandler(ctx *gin.Context) {
	var u User
	if err := ctx.BindJSON(&u); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	u.UUID = uuid.NewV4().String()
	listUser[u.UUID] = u
	ctx.JSON(http.StatusOK, u)
}
