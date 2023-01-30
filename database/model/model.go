package model

type User struct {
	ID       string `json:"id" gorm:"id"`
	Password string `json:"password" gorm:"password"`
	Type     int    `json:"type" gorm:"type"`
	Name     string `json:"name" gorm:"name"`
}

type KnowledgePoint struct {
	ID            int    `json:"id" gorm:"id"`
	Name          string `json:"name" gorm:"name"`
	FacilityValue int    `json:"facility_value" gorm:"facility_value"`
}

type Class struct {
	ID           int    `json:"id" gorm:"id"`
	Type         int    `json:"type" gorm:"type"`
	Name         string `json:"name" gorm:"name"`
	Introduction string `json:"introduction" gorm:"introduction"`
}

type Teach struct {
	ID        int    `json:"id" gorm:"id"`
	TeacherID string `json:"teacher_id" gorm:"teacher_id"`
	ClassID   int    `json:"class_id" gorm:"class_id"`
}

type Elective struct {
	ID        int    `json:"id" gorm:"id"`
	StudentID string `json:"student_id" gorm:"student_id"`
	ClassID   int    `json:"class_id" gorm:"class_id"`
}

type Subject struct {
	ID      int    `json:"id" gorm:"id"`
	Content string `json:"content" gorm:"content"`
	Answer  string `json:"answer" gorm:"answer"`
	Type    int    `json:"type" gorm:"type"`
}

type TestPaper struct {
	ID            int    `json:"id" gorm:"id"`
	SubjectList   string `json:"subject_list" gorm:"subject_list"`
	Creater       string `json:"creater" gorm:"creater"`
	FacilityValue int    `json:"facility_value" gorm:"facility_value"`
}
