package internal

import (
	"log"

	"github.com/meguriri/OnlineAssessmentSystem/database/model"
	database "github.com/meguriri/OnlineAssessmentSystem/database/operation"
)

func CreateClass(name string, typ int, introduction string) error {
	class := model.Class{
		Name:         name,
		Type:         typ,
		Introduction: introduction,
	}

	err := database.CreateClass(class)
	if err != nil {
		log.Printf("[CreateClass] CreateClass err,class:%+v,err:%+v", class, err)
		return err
	}

	return nil
}

func DeleteClass(id int) error {

	err := database.DeleteClass(id)
	if err != nil {
		log.Printf("[DeleteClass] DeleteClass err,id:%+v,err:%+v", id, err)
		return err
	}

	return nil
}

func UpdateClass(id int, name string, typ int, introduction string) error {
	class := model.Class{
		ID:           id,
		Name:         name,
		Type:         typ,
		Introduction: introduction,
	}

	err := database.UpdateClass(id, class)
	if err != nil {
		log.Printf("[UpdateClass] UpdateClass err,class:%+v,err:%+v", class, err)
		return err
	}

	return nil
}

func AddTeacher(teacherID string, classID int) error {

	err := database.AddTeach(teacherID, classID)
	if err != nil {
		log.Printf("[AddTeacher] AddElective err,teacherID:%+v,classID:%+v", teacherID, classID)
		return err
	}

	return nil
}

func Elective(studentID string, classID int) error {
	err := database.AddElective(studentID, classID)
	if err != nil {
		log.Printf("[Elective] AddElective err,studentID:%+v,classID:%+v", studentID, classID)
		return err
	}

	return nil
}
