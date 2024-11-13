package main

import (
	"delivery/api"
	"delivery/cmd"
	"github.com/gin-gonic/gin"
)

func main() {
	_ = cmd.InitDataBase()

	r := gin.Default()

	api.RoutesDefine(r)

	err := r.Run(":8080")
	if err != nil {
		return
	}
}
