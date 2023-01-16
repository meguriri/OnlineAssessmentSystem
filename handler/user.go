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

func Login(c *gin.Context) {
	user := body.UserLoginReq{}

	err := util.PostForm(c, &user)
	if err != nil {
		log.Printf("[Login] PostForm err,err:%+v", err)
		c.JSON(http.StatusOK, body.Res{
			Code: constant.CodeErr,
			Msg:  err.Error(),
		})
		return
	}

	err = internal.Login(user.ID, user.Password, user.Type)
	if err != nil {
		log.Printf("[Login] Login err,user:%+v,err:%+v", user, err)
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

func Register(c *gin.Context) {
	user := body.RegisterReq{}

	err := util.PostForm(c, &user)
	if err != nil {
		log.Printf("[Register] PostForm err,err:%+v", err)
		c.JSON(http.StatusOK, body.Res{
			Code: constant.CodeErr,
			Msg:  err.Error(),
		})
		return
	}

	if err := internal.CreateUser(user.ID, user.Name, user.Password, user.Type); err != nil {
		log.Printf("[Register] CreateUser err,user:%+v,err:%+v", user, err)
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
