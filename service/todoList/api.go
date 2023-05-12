package todolist

import (
	"AhZhengGo/service/todoList/handler"

	"github.com/gin-gonic/gin"
)

func UseApiGroup(r *gin.RouterGroup) {
	api := r.Group("/user/")

	api.GET("", handler.GetUser)
}
