package main

import (
	"github.com/gin-gonic/gin"
	"github.com/meguriri/OnlineAssessmentSystem/router"
)

func main() {
	r := gin.Default()
	router.InitRouter(r)

	r.Run(":8080")
}
