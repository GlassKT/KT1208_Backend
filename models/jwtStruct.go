package models

type TokenDetails struct {
	AccessToken string
	AccessUuid  string // redis에 저장할 때 key값
	AtExpires   int64
}

type AccessDetails struct { // redis에 접근하기 위한 구조체
	AccessUuid string
	UserId     string
}
