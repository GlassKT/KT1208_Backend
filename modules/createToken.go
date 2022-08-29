package modules

import (
	"log"
	"time"

	"github.com/dgrijalva/jwt-go"
)

var tokenString = []byte("secret")

func CreateToken(id string) (string, error) {

	// claims = 토큰의 조건
	claims := jwt.MapClaims{} // 조건을 맵형태로 저장

	claims["user_id"] = id
	claims["exp"] = time.Now().Add(time.Minute * 1).Unix()

	// SigningMethodHS256 형식으로(중요) claims를 참조한 엑세스 생성
	access := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	// 엑세스의 사인된 스트링형 반환
	token, err := access.SignedString([]byte(tokenString))

	if err != nil {
		log.Println("에러 메세지", err)
		return "", err
	}
	return token, nil
}

/*func VerfyToken(g *gin.Context) {

	//g.Get("user_token").(*jwt.Token)

	// 쿠키의 access-token 가져오기
	nToken, err := g.Request.Cookie("access-token")
	if err != nil {
		g.JSON(400, gin.H{
			"status":  "400",
			"message": "쿠기가 존재하지 않습니다",
		})
		g.Abort() // 두가지 요청 구분짓기
		return
	}
	// 쿠기의 토큰값 가져오기
	tokenValue := nToken.Value

	if tokenValue == "" {
		g.JSON(400, gin.H{
			"status":  "400",
			"message": "쿠기가 없습니다",
		})
		g.Abort()
		return
	}

	claims := jwt.MapClaims{}

	token, err := jwt.ParseWithClaims(tokenValue, &claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(tokenString), nil
	})

	if err != nil {
		g.JSON(400, gin.H{
			"status":  "400",
			"message": "토큰 인증 실패",
		})
	}

	fmt.Printf("token : %v", token)

	for key, val := range claims {
		fmt.Printf("key : %v, value : %v\n", key, val)
	}

	g.JSON(200, gin.H{
		"status":  "200",
		"message": "토큰 인증 성공",
	})
	return
}*/
