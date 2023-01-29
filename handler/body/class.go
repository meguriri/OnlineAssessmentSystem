package body

type CreateClassReq struct {
	Type         int    `json:"type"`
	Name         string `json:"name"`
	Introduction string `json:"introduction"`
}

type DeleteClassReq struct {
	ID int `json:"id"`
}

type UpdateClassReq struct {
	ID           int    `json:"id"`
	Type         int    `json:"type"`
	Name         string `json:"name"`
	Introduction string `json:"introduction"`
}

type AddTeacherReq struct {
	TeacherID string `json:"teacher_id"`
	ClassID   int    `json:"class_id"`
}

type ElectiveReq struct {
	StudentID string `json:"student_id"`
	ClassID   int    `json:"class_id"`
}
