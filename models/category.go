package models

import "time"

type Category struct {
	Id      int `orm:"pk;auto"`
	Name    string
	Created time.Time `orm:"auto_now_add;type(datetime);"`
	Updated time.Time `orm:"auto_now;type(datetime);"`
}
