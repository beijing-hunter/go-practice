package models

type User struct{
	Id int `gorm:"primary_key"`
	Username string
	Password string
}