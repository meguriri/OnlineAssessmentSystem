package database

import (
	"log"

	"github.com/meguriri/OnlineAssessmentSystem/constant"
	"github.com/meguriri/OnlineAssessmentSystem/database/model"
	"github.com/meguriri/OnlineAssessmentSystem/util"
)

func CreateTestPaper(paper model.TestPaper) error {
	if err := db.Table(constant.TableTestPaper).Create(&paper).Error; err != nil {
		log.Printf("[CreateTestPaper] Create err,paper:%+v,err:%+v", paper, err)
		return err
	}
	return nil
}

func GetTestPaperByID(id int) (model.TestPaper, error) {
	paper := model.TestPaper{}
	if err := db.Table(constant.TableTestPaper).Where("id=?", id).First(&paper).Error; !util.DBQueryErr(err) {
		log.Printf("[GetTestPaperByID] First err,id:%+v,err:%+v", id, err)
		return paper, err
	}

	return paper, nil
}
