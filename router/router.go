package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func InitRouter(router *gin.Engine) {
	router.LoadHTMLGlob("static/html/*")
	router.StaticFS("/css", http.Dir(".static/css"))
	router.StaticFS("/images", http.Dir(".static/images"))
	router.StaticFS("/js", http.Dir(".static/js"))
}