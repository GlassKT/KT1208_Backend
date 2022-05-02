package models

/*func sendMessage(g gin.Context){
		// 메일서버 로그인 정보 설정
		auth := smtp.PlainAuth("", "sender@live.com", "pwd", "smtp.live.com")

		from := "sender@live.com"
		to := []string{"sh-ym3384@naver.com"} // 복수 수신자 가능

		// 메시지 작성
		headerSubject := "Subject: 테스트\r\n"
		headerBlank := "\r\n"
		body := "메일 테스트입니다\r\n"
		msg := []byte(headerSubject + headerBlank + body)

		// 메일 보내기
		err := smtp.SendMail("smtp.live.com:587", auth, from, to, msg)
		if err != nil {
			panic(err)
		}
}*/
