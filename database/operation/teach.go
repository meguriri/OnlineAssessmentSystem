package database

import (
	"log"

	"github.com/meguriri/OnlineAssessmentSystem/constant"
	"github.com/meguriri/OnlineAssessmentSystem/database/model"
)

func AddTeach(teacherID string, classID int) error {
	teach := model.Teach{
		TeacherID: teacherID,
		ClassID:   classID,
	}

	if err := db.Table(constant.TableTeach).Create(&teach).Error; err != nil {
		log.Printf("[AddTeach] Create err,teach:%+v,err:%+v", teach, err)
		return err
	}

	return nil
}
