package model

type QA struct {
	Request  string `json:"request" bson:"request"`
	Response string `json:"response" bson:"response"`
}
