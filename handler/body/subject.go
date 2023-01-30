package body

type CreateSubjectReq struct {
	Content string `json:"content"`
	Answer  string `json:"answer"`
	Type    int    `json:"type"`
}

type UpdateSubjectReq struct {
	ID      int    `json:"id"`
	Content string `json:"content"`
	Answer  string `json:"answer"`
	Type    int    `json:"type"`
}

type DeleteSubjectReq struct {
	ID int `json:"id"`
}
