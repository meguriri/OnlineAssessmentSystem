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

func CreateClass(c *gin.Context) {
	class := body.CreateClassReq{}
	err := util.PostForm(c, &class)
	if err != nil {
		log.Printf("[CreateClass] PostForm err,err:%+v", err)
		c.JSON(http.StatusOK, body.Res{
			Code: constant.CodeErr,
			Msg:  err.Error(),
		})
		return
	}

	err = internal.CreateClass(class.Name, class.Type, class.Introduction)
	if err != nil {
		log.Printf("[CreateClass] CreateClass err,class:%+v,err:%+v", class, err)
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

func DeleteClass(c *gin.Context) {
	req := body.DeleteClassReq{}
	err := util.PostForm(c, &req)
	if err != nil {
		log.Printf("[UpdateClass] PostForm err,err:%+v", err)
		c.JSON(http.StatusOK, body.Res{
			Code: constant.CodeErr,
			Msg:  err.Error(),
		})
		return
	}

	err = internal.DeleteClass(req.ID)
	if err != nil {
		log.Printf("[DeleteClass] DeleteClass err,id:%+v,err:%+v", req.ID, err)
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

func UpdateClass(c *gin.Context) {
	class := body.UpdateClassReq{}
	err := util.PostForm(c, &class)
	if err != nil {
		log.Printf("[UpdateClass] PostForm err,err:%+v", err)
		c.JSON(http.StatusOK, body.Res{
			Code: constant.CodeErr,
			Msg:  err.Error(),
		})
		return
	}

	err = internal.UpdateClass(class.ID, class.Name, class.Type, class.Introduction)
	if err != nil {
		log.Printf("[UpdateClass] UpdateClass err,class:%+v,err:%+v", class, err)
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

func AddTeacherToClass(c *gin.Context) {
	addInfo := body.AddTeacherReq{}
	err := util.PostForm(c, &addInfo)
	if err != nil {
		log.Printf("[AddTeacherToClass] PostForm err,err:%+v", err)
		c.JSON(http.StatusOK, body.Res{
			Code: constant.CodeErr,
			Msg:  err.Error(),
		})
		return
	}

	err = internal.AddTeacher(addInfo.TeacherID, addInfo.ClassID)
	if err != nil {
		log.Printf("[AddTeacherToClass] AddTeacher err,info:%+v,err:%+v", addInfo, err)
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

func Elective(c *gin.Context) {
	elective := body.ElectiveReq{}
	err := util.PostForm(c, &elective)
	if err != nil {
		log.Printf("[Elective] PostForm err,err:%+v", err)
		c.JSON(http.StatusOK, body.Res{
			Code: constant.CodeErr,
			Msg:  err.Error(),
		})
		return
	}

	err = internal.Elective(elective.StudentID, elective.ClassID)
	if err != nil {
		log.Printf("[Elective] AddTeacher err,info:%+v,err:%+v", elective, err)
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
