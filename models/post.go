package models

import "time"

type Post struct {
	Id           int   `orm:"pk;auto"`
	User         *User `orm:"rel(fk)"`
	Title        string
	Url          string
	Content      string `orm:"type(text)"`
	Tags         string
	Views        int `orm:"default(0)"`
	IsTop        int8
	Created      time.Time `orm:"auto_now_add;type(datetime);"`
	Updated      time.Time `orm:"auto_now;type(datetime);"`
	CategoryId   int
	Status       int8
	Types        int8
	Info         string
	Image        string
	CommentCount int `orm:"default(0)"`
}
