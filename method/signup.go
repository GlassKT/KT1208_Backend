package method

import (
	_ "database/sql"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

// 받은 데이터의 ID가 DB에 존재한다면 string, false 반환
// 없다면 받은 데이터를 데이터베이스에 넣고 string, True를 반환

type signUp struct {
	ID        string    `gorm:"column:id" form:"id" binding:"required"`
	PW        string    `gorm:"column:pw" json:"pw" form:"pw" binding:"required"`
	NAME      string    `gorm:"column:name" json:"name" binding:"required"`
	EMAIL     string    `gprm:"column:email" json:"email"`
	CREATE_AT time.Time `gorm:"column:createAt`
}

func SignUp(g gin.Context) {

	user := new(signUp) // 새로운 유저객체

	err := g.Bind(user) // g의 값을 user에 파싱
	if err != nil {
		g.JSON(http.StatusBadRequest, err.Error())
	}

}
