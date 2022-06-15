package database

import (
	"time"
)

// DB의 계정의 모든 정보

type User struct {
	ID    string `gorm:"primary_key; column:id" json:"id" form:"id"`
	PW    string `gorm:"column:pw" json:"pw,omitempty" form:"pw"` //,omitempty = null값이면 그냥 없다치고 업로드
	NAME  string `gorm:"column:name" json:"name" form:"name"`
	EMAIL string `gprm:"column:email" json:"email" form:"email"`
	//CREATE_AT time.Time `gorm:"column:createAt;autoCreateTime" json:"createAt,omitempty"`
	CREATE_AT time.Time `gorm:"column:createAt" json:"createAt,omitempty"`

	//BIRTH     time.Time    `gorm:"column:birth" json:"birth" form:"birth"`
	BIRTH     time.Time    `gorm:"column:birth" json:"birth" form:"birth" time_format:"RFC3339"`
	MBTI      string       `gorm:"column:mbti" json:"mbti" form:"mbti"`
	AREA      string       `gorm:"column:area" json:"area" form:"area"`
	SCHOOL    string       `gorm:"column:school" json:"school" form:"school"`
	INTRODUCE string       `gorm:"column:introduce" json:"introduce" form:"introduce"`
	IMGNAME   string       `gorm:"column:imgname" json:"imgname" form:"imgname"`
	FRIENDS   []Makefriend `gorm:"foreignKey:USERID;association_foreignkey:ID"`

	HobbyId int   `gorm:"column:hobbyId" json:"hobbyId" form:"hobbyId"`
	Hobby   Hobby `gorm:"foreignKey:HobbyId"`
}

type Makefriend struct {
	USERID   string `gorm:"column:user_id" json:"user_id" form:"user_id"`
	FRIENDID string `gorm:"column:friend_id" json:"friend_id" form:"friend_id"`
	STATUS   string `gorm:"column:status" json:"status" form:"status"`
}

type Hobby struct {
	CODE int    `gorm:"primary_key; column:code" json:"code" form:"code"`
	NAME string `gorm:"column:name" json:"name" form:"name"`
}
