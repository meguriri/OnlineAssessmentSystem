package internal

import (
	"errors"
	"log"

	"github.com/meguriri/OnlineAssessmentSystem/constant"
	"github.com/meguriri/OnlineAssessmentSystem/database/model"
	database "github.com/meguriri/OnlineAssessmentSystem/database/operation"
)

func Login(id string, password string, typ int) error {
	user, err := database.GetUserByID(id)
	if err != nil {
		log.Printf("[Login] GetUserByID err,id:%+v,err:%+v", id, err)
		return err
	}

	if user.Password != password || typ != user.Type {
		log.Printf("[Login] wrong password,id:%+v,user:%+v", id, user)
		return errors.New(constant.WrongIDOrPassword)
	}

	return nil
}

func CreateUser(id, name, password string, typ int) error {
	user := model.User{
		ID:       id,
		Name:     name,
		Password: password,
		Type:     typ,
	}

	err := database.CreateUser(user)
	if err != nil {
		log.Printf("[CreateUser] CreateUser err,user:%+v,err:%+v", user, err)
		return err
	}
	return nil
}

func UpdateUser(id, oldPassword, newPassword string) error {
	user, err := database.GetUserByID(id)
	if err != nil {
		log.Printf("[UpdateUser] GetUserByID err,id:%+v,err:%+v", id, err)
		return err
	}

	if user.Password != oldPassword {
		log.Printf("[UpdateUser] wrong old password,id:%+v,user:%+v", id, user)
		return errors.New(constant.WrongOldPassword)
	}

	user.Password = newPassword

	err = database.UpdateUserByID(user.ID, user)
	if err != nil {
		log.Printf("[UpdateUser] UpdateUserByID err,user:%+v,err:%+v", user, err)
		return err
	}

	return nil
}
