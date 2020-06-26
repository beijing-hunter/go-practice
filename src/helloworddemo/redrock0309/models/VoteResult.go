package models

type Voteresult struct{
	Id int `gorm:"primary_key"`
	Voteid int
	Xsuserid int
	Xsusername string
	Tpuserid int
	Tpusername string
}