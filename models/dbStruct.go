package models

import (
	"time"
)

// DB의 계정의 모든 정보

type User struct {
	Id        string `gorm:"primary_key; size:32" json:"id" form:"id"`
	Pw        string `gorm:"size:128" json:"pw" form:"pw"`
	Name      string `gorm:"size:32" json:"name" form:"name"`
	Email     string `gorm:"size:64" json:"email" form:"email"`
	Mbti      string `gorm:"size:64" json:"mbti" form:"mbti"`
	Area      string `gorm:"size:64" json:"area" form:"area"`
	School    string `gorm:"size:64" json:"school" form:"school"`
	Introduce string `gorm:"size:256" json:"introduce" form:"introduce"`
	ImgName   string `gorm:"size:64" json:"imgname" form:"imgname"`
	Birthday  string `gorm:"size:32" json:"birthday" form:"birthday"`

	Friends       []Friend       `gorm:"foreignkey:user"`
	BlockFriends  []BlockFriend  `gorm:"foreignkey:user"`
	Hobbys        []Hobby        `gorm:"foreignkey:user"`
	ChattingRooms []ChattingRoom `gorm:"foreignkey:owner"`

	CreatedAt time.Time
	//Birthday  time.Time `json:"birth" form:"birth" time_format:"RFC3339"`
}

type Friend struct {
	User   string `gorm:"size:32" json:"user"`
	Friend string `gorm:"size:32" json:"friend"`
	Status string `gorm:"size:16" json:"status"`
}

type BlockFriend struct {
	User   string `gorm:"size:32" json:"user"`
	Friend string `gorm:"size:32" json:"friend"`
}

type Hobby struct {
	User  string `gorm:"size:32" json:"user"`
	Hobby string `gorm:"size:64" json:"hobby"`
}

type ChattingRoom struct {
	RoomNum  int       `gorm:"primary_key;auto_increment"`
	Owner    string    `gorm:"size:32" json:"owner"`
	Guest    string    `gorm:"size:32" json:"guest"`
	Messages []Message `gorm:"foreignkey:RoomNum"`
}

type Message struct {
	RoomNum   int    `gorm:"size:64"`
	Content   string `json:"content"`
	User      string `gorm:"size:32" json:"user"`
	CreatedAt time.Time
}
