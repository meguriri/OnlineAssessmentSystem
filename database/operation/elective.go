package database

import (
	"log"

	"github.com/meguriri/OnlineAssessmentSystem/constant"
	"github.com/meguriri/OnlineAssessmentSystem/database/model"
)

func AddElective(studentID string, classID int) error {
	elective := model.Elective{
		StudentID: studentID,
		ClassID:   classID,
	}

	if err := db.Table(constant.TableElective).Create(&elective).Error; err != nil {
		log.Printf("[AddElective] Create err,elective:%+v,err:%+v", elective, err)
		return err
	}
	return nil
}
