package handler

import (
	"AhZhengGo/app/common/response"
	"net/http"

	"github.com/gin-gonic/gin"
)

type User struct {
	Name string
	Age  int
}

func GetUser(c *gin.Context) {
	res := User{
		Name: "Jerry",
		Age:  28,
	}

	c.JSON(http.StatusOK, response.OK.WithData(res))
}
