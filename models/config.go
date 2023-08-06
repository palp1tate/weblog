package models

type Config struct {
	Id    int `orm:"pk;auto"`
	Name  string
	Value string
}
