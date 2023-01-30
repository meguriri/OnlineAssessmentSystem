package handler

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/meguriri/OnlineAssessmentSystem/constant"
	"github.com/meguriri/OnlineAssessmentSystem/handler/body"
	"github.com/meguriri/OnlineAssessmentSystem/internal"
	"github.com/meguriri/OnlineAssessmentSystem/util"
)

func CreateSubject(c *gin.Context) {
	subject := body.CreateSubjectReq{}

	err := util.PostForm(c, &subject)
	if err != nil {
		log.Printf("[CreateSubject] PostForm err,err:%+v", err)
		c.JSON(http.StatusOK, body.Res{
			Code: constant.CodeErr,
			Msg:  err.Error(),
		})
		return
	}

	err = internal.CreateSubject(subject.Content, subject.Answer, subject.Type)
	if err != nil {
		log.Printf("[CreateSubject] CreateSubject err,subject:%+v,err:%+v", subject, err)
		c.JSON(http.StatusOK, body.Res{
			Code: constant.CodeErr,
			Msg:  err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, body.Res{
		Code: constant.CodeSuccess,
	})
}

func UpdateSubject(c *gin.Context) {
	subject := body.UpdateSubjectReq{}

	err := util.PostForm(c, &subject)
	if err != nil {
		log.Printf("[UpdateSubject] PostForm err,err:%+v", err)
		c.JSON(http.StatusOK, body.Res{
			Code: constant.CodeErr,
			Msg:  err.Error(),
		})
		return
	}

	err = internal.UpdateSubject(subject.ID, subject.Content, subject.Answer, subject.Type)
	if err != nil {
		log.Printf("[UpdateSubject] UpdateSubject err,subject:%+v,err:%+v", subject, err)
		c.JSON(http.StatusOK, body.Res{
			Code: constant.CodeErr,
			Msg:  err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, body.Res{
		Code: constant.CodeSuccess,
	})
}

func DeleteSubject(c *gin.Context) {
	req := body.DeleteSubjectReq{}

	err := util.PostForm(c, &req)
	if err != nil {
		log.Printf("[DeleteSubject] PostForm err,err:%+v", err)
		c.JSON(http.StatusOK, body.Res{
			Code: constant.CodeErr,
			Msg:  err.Error(),
		})
		return
	}

	err = internal.DeleteSubject(req.ID)
	if err != nil {
		log.Printf("[DeleteSubject] DeleteSubject err,req:%+v,err:%+v", req, err)
		c.JSON(http.StatusOK, body.Res{
			Code: constant.CodeErr,
			Msg:  err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, body.Res{
		Code: constant.CodeSuccess,
	})
}
