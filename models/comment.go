package models

import "time"

type Comment struct {
	Id       int `orm:"pk;auto"`
	Username string
	Content  string
	Created  time.Time `orm:"auto_now_add;type(datetime);"`
	PostId   int
	Ip       string
}
