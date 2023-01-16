package database

import (
	"log"

	"github.com/meguriri/OnlineAssessmentSystem/constant"
	"github.com/meguriri/OnlineAssessmentSystem/database/model"
	"github.com/meguriri/OnlineAssessmentSystem/util"
)

func GetUserByID(id string) (model.User, error) {
	user := model.User{}

	if err := db.Table(constant.TableUser).Where("id=?", id).First(&user).Error; !util.DBQueryErr(err) {
		log.Printf("[GetUserByID] First err,id:%+v,err:%+v", id, err)
		return user, err
	}

	return user, nil
}

func UpdateUserByID(id string, user model.User) error {
	if err := db.Table(constant.TableUser).Where("id=?", id).UpdateColumns(&user).Error; err != nil {
		log.Printf("[UpdateUserByID] UpdateColumns err,id:%+v,user:%+v", id, user)
		return err
	}

	return nil
}

func CreateUser(user model.User) error {
	if err := db.Table(constant.TableUser).Create(&user).Error; err != nil {
		log.Printf("[CreateUser] Create err,user:%+v,err:%+v", user, err)
		return err
	}
	return nil
}
