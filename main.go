package main

import (
	todolist "AhZhengGo/service/todoList"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	api := r.Group("/api/")

	todolist.UseApiGroup(api)

	err := r.Run(":8081")

	if err != nil {
		log.Fatal(err)
	}
}
