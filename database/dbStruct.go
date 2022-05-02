package database

// DB의 계정의 모든 정보

type User struct {
	ID    string `gorm:"primary_key; column:id" json:"id" form:"id"`
	PW    string `gorm:"column:pw" json:"pw" form:"pw"`
	NAME  string `gorm:"column:name" json:"name" form:"name"`
	EMAIL string `gprm:"column:email" json:"email" form:"email"`
	//CREATE_AT time.Time `gorm:"column:createAt;autoCreateTime" json:"createAt"`

	//BIRTH     string       `gorm:"column:birth; sql:DEFAULT:0000-00-00" json:"birth" form:"birth"`
	MBTI      string       `gorm:"column:mbti" json:"mbti" form:"mbti"`
	AREA      string       `gorm:"column:area" json:"area" form:"area"`
	SCHOOL    string       `gorm:"column:school" json:"school" form:"school"`
	INTRODUCE string       `gorm:"column:introduce" json:"introduce" form:"introduce"`
	IMGNAME   string       `gorm:"column:imgname" json:"imgname" form:"imgname"`
	FRIENDS   []Makefriend `gorm:"foreignKey:USERID;association_foreignkey:ID"`
}

type Makefriend struct {
	USERID   string `gorm:"column:user_id" json:"user_id" form:"user_id"`
	FRIENDID string `gorm:"column:friend_id" json:"friend_id" form:"friend_id"`
}
