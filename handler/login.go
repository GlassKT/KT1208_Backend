package method

import (
	db "GlassKT/database"
	"GlassKT/method"

	"github.com/gin-gonic/gin"
)

type loginstruct struct {
	ID string `form :"id" json:"id"`
	PW string `form :"pw" json:"pw"`
}

func Login(g *gin.Context) {
	param := new(loginstruct)
	err := g.Bind(param) // g의 값을 user에 파싱
	if err != nil {
		g.JSON(400, gin.H{"message": "파싱중 오류"})
		return
	}

	User := &db.User{} // orm에서 원본 DB 접근을 위한 객체

	if param.ID == "" || param.PW == "" {
		g.JSON(400, gin.H{"message": "ID,PW,NAME 모두 입력하세요"})
		return
	}

	if err := db.DB.Where("id = ? AND pw = ?", param.ID, param.PW).Find(User).Error; err != nil {
		g.JSON(400, gin.H{"message": "ID가 존재하지 않거나 PW가 잘못되었습니다"})
		return
	}

	token, err := method.CreateToken(param.ID)
	if err != nil {
		g.JSON(400, gin.H{"message": "토큰 생성중 오류"})
		return
	}
	g.JSON(200, token)
	return

}
