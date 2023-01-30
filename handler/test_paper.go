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

func CreateTestPaper(c *gin.Context) {
	req := body.CreateTestPaperReq{}

	err := util.PostForm(c, &req)
	if err != nil {
		log.Printf("[CreateTestPaper] PostForm err,err:%+v", err)
		c.JSON(http.StatusOK, body.Res{
			Code: constant.CodeErr,
			Msg:  err.Error(),
		})
		return
	}

	testPaper, err := internal.CreateTestPaper(req.UserID, req.FacilityValue)
	if err != nil {
		log.Printf("[CreateTestPaper] CreateTestPaper err,req:%+v,err:%+v", req, err)
		c.JSON(http.StatusOK, body.Res{
			Code: constant.CodeErr,
			Msg:  err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, body.Res{
		Code: constant.CodeSuccess,
		Data: testPaper,
	})
}
