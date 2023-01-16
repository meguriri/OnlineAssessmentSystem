package util

import (
	"encoding/json"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/meguriri/OnlineAssessmentSystem/constant"
)

func DBQueryErr(err error) bool {
	if err == nil || err.Error() == constant.ErrRecordNotFound {
		return true
	}
	return false
}

func PostForm(c *gin.Context, body interface{}) error {
	parms, err := c.GetRawData()
	if err != nil {
		log.Printf("[PostForm] GetRawData err,reqBody:%+v,err:%+v", c.Request.Body, err)
		return err
	}

	err = json.Unmarshal(parms, body)
	if err != nil {
		log.Printf("[PostForm] Unmarshal err,parms:%+v,err:%+v", parms, err)
		return err
	}
	return nil
}
