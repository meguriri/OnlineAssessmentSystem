package internal

import (
	"log"

	"github.com/meguriri/OnlineAssessmentSystem/database/model"
	database "github.com/meguriri/OnlineAssessmentSystem/database/operation"
)

func CreateSubject(content string, anwser string, typ int) error {
	subject := model.Subject{
		Content: content,
		Answer:  anwser,
		Type:    typ,
	}

	err := database.CreateSubject(subject)
	if err != nil {
		log.Printf("[CreateSubject] CreateSubject err,subject:%+v,err:%+v", subject, err)
		return err
	}

	return nil
}

func UpdateSubject(id int, content string, anwser string, typ int) error {
	subject := model.Subject{
		ID:      id,
		Content: content,
		Answer:  anwser,
		Type:    typ,
	}

	err := database.UpdateSubject(id, subject)
	if err != nil {
		log.Printf("[UpdateSubject] UpdateSubject err,id:%+v,subject:%+v", id, subject)
		return err
	}

	return nil
}

func DeleteSubject(id int) error {
	err := database.DeleteSubject(id)
	if err != nil {
		log.Printf("[DeleteSubject] DeleteSubject err,id:%+v,err:%+v", id, err)
		return err
	}

	return nil
}
