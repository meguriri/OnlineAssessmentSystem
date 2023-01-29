package router

import (
	"github.com/gin-gonic/gin"
	"github.com/meguriri/OnlineAssessmentSystem/handler"
)

func InitRouter(router *gin.Engine) {
	router.Static("/static", "./static/")
	router.LoadHTMLGlob("template/*")
	///////////////////////////////////////////////////////////////////////////////////////
	router.POST("/login", handler.Login)
	router.POST("/register", handler.Register)
	router.POST("/class/create", handler.CreateClass)
	router.POST("/class/delete", handler.DeleteClass)
	router.POST("/class/Update", handler.UpdateClass)
	router.POST("/class/add_teacher", handler.AddTeacherToClass)
	router.POST("/class/elective", handler.Elective)
	//////////////////////////////////////////////////////////////////////////////////////
	router.GET("/login", handler.GetLogin)
	router.GET("/", handler.GetLogin)
	router.POST("/user/update", handler.UpdateUser)
}
