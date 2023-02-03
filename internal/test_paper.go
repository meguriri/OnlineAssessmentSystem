package internal

import (
	"fmt"
	"log"
	"math/rand"
	"time"

	"github.com/meguriri/OnlineAssessmentSystem/constant"
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
			ID:      item.ID,
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

func GudgeAnwser(anwserList []body.Anwser) (body.SubmitAnwserRes, error) {
	gudgeAnwser := body.SubmitAnwserRes{}

	idList := []int{}
	for _, item := range anwserList {
		idList = append(idList, item.ID)
	}

	subjectList, err := database.GetSubject(idList)
	if err != nil {
		log.Printf("[GudgeAnwser] GetSubject err,idList:%+v,err:%+v", idList, err)
		return gudgeAnwser, err
	}

	subjectMap := map[int]model.Subject{}

	for _, item := range subjectList {
		subjectMap[item.ID] = item
	}

	for _, item := range anwserList {
		var typ int
		if item.Anwser != subjectMap[item.ID].Answer {
			typ = constant.WrongAnwser
		} else {
			typ = constant.TrueAnwser
		}

		gudgeAnwser.GudgeAnwser = append(gudgeAnwser.GudgeAnwser, body.GudgeAnwser{
			ID:         item.ID,
			TrueAnwser: subjectMap[item.ID].Answer,
			Content:    subjectMap[item.ID].Content,
			UserAnwser: item.Anwser,
			Type:       typ,
		})
	}

	return gudgeAnwser, nil
}
