package models

type Voteuser struct{
	Id int `gorm:"primary_key"`
	Voteid int
	Xsuserid int
	Xsusername string
	Votetotalcount int
}