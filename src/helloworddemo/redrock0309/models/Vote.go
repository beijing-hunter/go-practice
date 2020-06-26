package models

type Vote struct{
	Id int `gorm:"primary_key"`
	Vname string
	Starttime string
	Endtime string
}