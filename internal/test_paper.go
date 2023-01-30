package internal

import (
	"fmt"
	"log"
	"math/rand"
	"time"

	"github.com/meguriri/OnlineAssessmentSystem/database/model"
	database "github.com/meguriri/OnlineAssessmentSystem/database/operation"
	"github.com/meguriri/OnlineAssessmentSystem/handler/body"
)

func CreateTestPaper(userID string, facilityValue int) (body.TestPaperRes, error) {
	testPaper := body.TestPaperRes{
		Creater: userID,
	}

	maxId, err := database.GetSubjectMaxIncrementID()
	if err != nil {
		log.Printf("[CreateTestPaper] GetSubjectMaxIncrementID err,err:%+v", err)
		return testPaper, err
	}

	rand.Seed(time.Now().UnixNano())
	idList := []int{}
	for i := 0; i < 5; i++ {
		idList = append(idList, rand.Intn(maxId))
	}

	sujectList, err := database.GetSubject(idList)
	if err != nil {
		log.Printf("[CreateTestPaper] GetSubject err,idList:%+v,err:%+v", idList, err)
		return testPaper, err
	}

	paper := model.TestPaper{
		Creater:       userID,
		FacilityValue: facilityValue,
		SubjectList:   "",
	}

	for _, item := range sujectList {
		paper.SubjectList = fmt.Sprintf("%d,", item.ID)
		testPaper.SubjectList = append(testPaper.SubjectList, body.Subject{
			Content: item.Content,
			Type:    item.Type,
		})
	}

	paper.SubjectList = paper.SubjectList[:len(paper.SubjectList)-2] //去掉最后一个逗号

	err = database.CreateTestPaper(paper)
	if err != nil {
		log.Printf("[CreateTestPaper] CreateTestPaper err,paper:%+v,err:%+v", paper, err)
		return testPaper, err
	}

	return testPaper, nil
}
