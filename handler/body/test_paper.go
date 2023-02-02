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
	ID      int    `json:"id"`
	Content string `json:"content"`
	Type    int    `json:"type"`
}

type SubmitAnwserReq struct {
	AnwserList []Anwser `json:"anwser_list"`
}

type Anwser struct {
	ID     int    `json:"id"`
	Anwser string `json:"anwser"`
}

type SubmitAnwserRes struct {
	GudgeAnwser []GudgeAnwser `json:"gudge_anwser"`
}

type GudgeAnwser struct {
	ID         int    `json:"id"`
	Content    string `json:"content"`
	TrueAnwser string `json:"true_anwser"`
	UserAnwser string `json:"user_anwser"`
	Type       int    `json:"type"`
}
