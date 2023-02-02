package database

import (
	"log"

	"github.com/meguriri/OnlineAssessmentSystem/constant"
	"github.com/meguriri/OnlineAssessmentSystem/database/model"
	"github.com/meguriri/OnlineAssessmentSystem/util"
)

func CreateSubject(subject model.Subject) error {
	if err := db.Table(constant.TableSubject).Create(&subject).Error; err != nil {
		log.Printf("[CreateSubject] Create err,subject:%+v,err:%+v", subject, err)
		return err
	}
	return nil
}

func GetSubjectMaxIncrementID() (int, error) {
	var maxID int
	if err := db.Table(constant.TableSubject).Select("max(id)").Row().Scan(&maxID); err != nil {
		log.Printf("[GetSubjectMaxIncrementID] Scan err,err:%+v", err)
		return 0, err
	}

	return maxID, nil
}

func GetAllSubject() ([]model.Subject, error) {
	subjectList := []model.Subject{}

	if err := db.Table(constant.TableSubject).Find(&subjectList).Error; !util.DBQueryErr(err) {
		log.Printf("[GetAllSubject] Find err,err:%+v", err)
		return nil, err
	}
	return subjectList, nil
}

func GetSubject(id []int) ([]model.Subject, error) {
	subjectList := []model.Subject{}

	if err := db.Table(constant.TableSubject).Where("id in (?)", id).Find(&subjectList).Error; !util.DBQueryErr(err) {
		log.Printf("[GetSubject] Find err,idlist:%+v,err:%+v", id, err)
		return nil, err
	}
	return subjectList, nil
}

func UpdateSubject(id int, subject model.Subject) error {
	if err := db.Table(constant.TableSubject).Where("id=?", id).UpdateColumns(subject).Error; err != nil {
		log.Printf("[UpdateSubject] UpdateColumns err,id:%+v,subject:%+v,err:%+v", id, subject, err)
		return err
	}
	return nil
}

func DeleteSubject(id int) error {
	if err := db.Table(constant.TableSubject).Delete(model.Subject{ID: id}).Error; err != nil {
		log.Printf("[DeleteSubject] Delete err,id:%+v,err:%+v", id, err)
		return err
	}
	return nil
}
