package models

import (
	"time"
)

type User struct {
	Id         int `orm:"pk;auto"`
	Username   string
	Password   string
	Email      string
	LoginCount int `orm:"default(0)"`
	LastTime   time.Time
	LastIp     string
	State      int8
	Created    time.Time `orm:"auto_now_add;type(datetime);"`
	Updated    time.Time `orm:"auto_now;type(datetime);"`
}
