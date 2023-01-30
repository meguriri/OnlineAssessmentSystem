package body

type CreateTestPaperReq struct {
	UserID        string `json:"user_id"`
	FacilityValue int    `json:"facility_value"`
}

type TestPaperRes struct {
	Creater     string    `json:"creater"`
	SubjectList []Subject `json:"subject_list"`
}

type Subject struct {
	Content string `json:"content"`
	Type    int    `json:"type"`
}
