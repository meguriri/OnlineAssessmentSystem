package main

import (
	"github.com/gin-gonic/gin"
	database "github.com/meguriri/OnlineAssessmentSystem/database/operation"
	"github.com/meguriri/OnlineAssessmentSystem/router"
)

func main() {
	r := gin.Default()

	database.InitDB()
	router.InitRouter(r)

	r.Run(":8080")
}
