package model

import (
	"fmt"
	"strconv"
)

type Follower struct {
	ID         uint
	UserID     uint `gorm:"not null" sql:"index"`
	FollowerID uint `gorm:"not null" sql:"index"`
}

func FollowUser(userid uint, followersid ...uint) error {
	sqlstr := "insert into followers (user_id,follower_id) values "
	for k, v := range followersid {
		sqlstr += fmt.Sprintf("(%d,%d)", userid, v)
		if k < len(followersid)-1 {
			sqlstr += ","
		}
	}
	return DB.Exec(sqlstr).Error
}

func UnfollowUser(userid uint, followersid ...uint) error {
	sqlstr := fmt.Sprintf("delete from followers where follower_id=%d and user_id in (", userid)
	for k, v := range followersid {
		sqlstr += strconv.Itoa(int(v))
		if k < len(followersid)-1 {
			sqlstr += ","
		}
	}
	sqlstr += ")"
	return DB.Exec(sqlstr).Error
}

func ListFollowers() {}