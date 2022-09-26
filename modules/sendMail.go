package modules

import (
	"fmt"

	"gopkg.in/gomail.v2"
)

func SendMail(authNum, toEmail string) error {

	m := gomail.NewMessage()
	m.SetHeader("From", "sh-ym3384@naver.com")
	m.SetHeader("To", toEmail, "sh-ym3384@naver.com")
	m.SetAddressHeader("Cc", "sh-ym3384@naver.com", "name")
	m.SetHeader("Subject", "email test")
	m.SetBody("text/html", authNum)

	d := gomail.NewDialer("smtp.naver.com", 465, "sh-ym3384@naver.com", "TAKIYOU^^123")

	if err := d.DialAndSend(m); err != nil {
		fmt.Println(err)
		return err
	}

	return nil
}
