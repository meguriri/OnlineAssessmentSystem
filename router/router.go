package router

import (
	"github.com/gin-gonic/gin"
	"github.com/meguriri/OnlineAssessmentSystem/handler"
)

func InitRouter(router *gin.Engine) {
	router.Static("/static", "./static/")
	router.LoadHTMLGlob("template/*")
	router.POST("/login", handler.Login)
	router.POST("/register", handler.Register)
}
