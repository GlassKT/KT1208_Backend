package models

import (
	"time"
)

// DB의 계정의 모든 정보

type User struct {
	Id        string `gorm:"primaryKey; size:32" json:"id" form:"id"`
	Email     string `gorm:"size:64" json:"email" form:"email"`
	Name      string `gorm:"size:32" json:"name" form:"name"`
	Mbti      string `gorm:"size:64" json:"mbti" form:"mbti"`
	Area      string `gorm:"size:64" json:"area" form:"area"`
	School    string `gorm:"size:64" json:"school" form:"school"`
	Introduce string `gorm:"size:256" json:"introd$uce" form:"introduce"`
	Birthday  string `gorm:"size:32" json:"birthday" form:"birthday"`

	Pw           string        `gorm:"size:128" json:"pw" form:"pw"`
	ImgName      string        `gorm:"size:64" json:"imgname" form:"imgname"`
	Friends      []Friend      `gorm:"foreignkey:user"`
	BlockFriends []BlockFriend `gorm:"foreignkey:user"`
	Hobbys       []Hobby       `gorm:"foreignkey:user"`

	CreatedAt time.Time
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
	User  string `gorm:"size:32"`
	Hobby string `gorm:"size:64" json:"hobby"`
}

type ChattingRoom struct {
	User    string `gorm:"size:32" json:"id"`
	RoomNum string `gorm:"size:256" json:"roomnum"`

	Messages []Message `gorm:"foreignKey:NumRoom;references:RoomNum"`
}

type Message struct {
	NumRoom  string `gorm:"size:64" json:"roomnum"`
	Content  string `gorm:"size:256" json:"content"`
	UserId   string `gorm:"size:32" json:"id"`
	Name     string `gorm:"size:32" json:"name"`
	Createat string `gorm:"size:64" json:"date"`

	Messagescol int
}
