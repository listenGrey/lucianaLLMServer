package model

type QA struct {
	Request  string `json:"request" bson:"request"`
	Response string `json:"response" bson:"response"`
}

type Request struct {
	Cid    int64  `json:"cid" binding:"required"`
	Prompt string `json:"prompt" binding:"required"`
}
