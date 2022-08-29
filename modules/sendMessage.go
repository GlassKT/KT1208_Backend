package modules

import (
	"log"

	"gopkg.in/gomail.v2"
)

func SendMail(randomnum, toMail string) error {

	m := gomail.NewMessage()
	m.SetHeader("From", "glasskt@naver.com")
	m.SetHeader("To", toMail, "glasskt@naver.com")
	m.SetAddressHeader("Cc", "glasskt@naver.com", "name")
	m.SetHeader("Subject", "email test")
	m.SetBody("text/html", randomnum)

	d := gomail.NewDialer("smtp.naver.com", 465, "glasskt", "glassKT112!")

	if err := d.DialAndSend(m); err != nil {
		return err
	}

	log.Print("[메일보내기 성공]")

	return nil
}
