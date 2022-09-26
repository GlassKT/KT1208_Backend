package models

type SignupInfo struct {
	Id    string `json:"id"`
	Pw    string `json:"pw"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

type LoginInfo struct {
	Id string `json:"id"`
	Pw string `json:"pw"`
}

type EditInfo struct {
	Id        string `json:"id"`
	Pw        string `json:"pw"`
	Name      string `json:"name"`
	Email     string `json:"email"`
	Mbti      string `json:"mbti"`
	Area      string `json:"area"`
	School    string `json:"school"`
	Introduce string `json:"introduce"`
	Hobby     string `json:"hobby"`
	Birthday  string `json:"birthday"`
}

type WebFriendInfo struct {
	User   string `json:"user"`
	Friend string `json:"friend"`
}

type WebFriendHobbyInfo struct {
	User   string   `json:"user"`
	Friend string   `json:"friend"`
	Hobby  []string `json:"hobby"`
}

type WebHobbyInfo struct {
	User  string   `json:"user"`
	Hobby []string `json:"hobby"`
}

/*type GetFriendInfo struct {
	Id        string   `json:"id"`
	Name      string   `json:"name"`
	Mbti      string   `json:"mbti"`
	Area      string   `json:"area"`
	School    string   `json:"school"`
	Introduce string   `json:"introduce"`
	Hobby     []string `json:"hobby"`
	Birthday  string   `json:"birthday"`
}*/

type GetFriendInfo struct {
	Id    string   `json:"id"`
	Name  string   `json:"name"`
	Hobby []string `json:"hobby"`
}

type GetUserInfo struct {
	Id        string   `json:"id"`
	Name      string   `json:"name"`
	Mbti      string   `json:"mbti"`
	Area      string   `json:"area"`
	School    string   `json:"school"`
	Introduce string   `json:"introduce"`
	Hobby     []string `json:"hobby"`
}

type WebUserInfo struct {
	Id    string   `json:"id"`
	Hobby []string `json:"hobby"`
}
