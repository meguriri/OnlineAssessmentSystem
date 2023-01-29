package database

import (
	"log"

	"github.com/meguriri/OnlineAssessmentSystem/constant"
	"github.com/meguriri/OnlineAssessmentSystem/database/model"
)

func GetClassList() ([]model.Class, error) {
	classList := []model.Class{}

	return classList, nil
}

func GetClass(id string) (model.Class, error) {
	class := model.Class{}

	return class, nil
}

func CreateClass(class model.Class) error {
	if err := db.Table(constant.TableClass).Create(&class).Error; err != nil {
		log.Printf("[CreateClass] Create err,class:%+v,err:%+v", class, err)
		return err
	}

	return nil
}

func UpdateClass(id int, class model.Class) error {
	if err := db.Table(constant.TableClass).UpdateColumns(class).Error; err != nil {
		log.Printf("[UpdateClass] UpdateColumns err,id:%+v,class:%+v,err:%+v", id, class, err)
		return err
	}

	return nil
}

func DeleteClass(id int) error {
	if err := db.Table(constant.TableClass).Delete(&model.Class{ID: id}).Error; err != nil {
		log.Printf("[DeleteClass] Delete err,id:%+v,err:%+v", id, err)
		return err
	}

	return nil
}
