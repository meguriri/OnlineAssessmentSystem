package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetLogin(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", gin.H{})
}

func GetClass(c *gin.Context) {
	c.HTML(http.StatusOK, "class.html", gin.H{})
}

func GetHome(c *gin.Context) {
	c.HTML(http.StatusOK, "home.html", gin.H{})
}

func AddClass(c *gin.Context) {
	c.HTML(http.StatusOK, "add_class.html", gin.H{})
}

func GetSubject(c *gin.Context) {
	c.HTML(http.StatusOK, "subject.html", gin.H{})
}

func GetAddSubject(c *gin.Context) {
	c.HTML(http.StatusOK, "add_subject.html", gin.H{})
}

func GetTestPaper(c *gin.Context) {
	c.HTML(http.StatusOK, "test_paper.html", gin.H{})
}
