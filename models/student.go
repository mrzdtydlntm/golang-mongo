package models

type Student struct {
	Name  string `bson:"name"`
	Grade int    `bson:"Grade"`
}
