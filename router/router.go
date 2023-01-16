package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/meguriri/OnlineAssessmentSystem/handler"
)

func InitRouter(router *gin.Engine) {
	router.LoadHTMLGlob("static/html/*")
	router.StaticFS("/css", http.Dir(".static/css"))
	router.StaticFS("/images", http.Dir(".static/images"))
	router.StaticFS("/js", http.Dir(".static/js"))
	router.POST("/login", handler.Login)
	router.POST("/register", handler.Register)
}
