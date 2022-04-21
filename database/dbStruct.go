package database

import "time"

// DB의 계정의 모든 정보

type User struct {
	ID        string    `gorm:"column:id" json:"id"`
	PW        string    `gorm:"column:pw" json:"pw"`
	NAME      string    `gorm:"column:name" json:"name"`
	EMAIL     string    `gprm:"column:email" json:"email"`
	CREATE_AT time.Time `gorm:"column:createAt`
}
