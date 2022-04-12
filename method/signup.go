package signin

type SignUpParam struct {
	ID    string `json:"id" form:"id" query:"id"`
	PW    string `json:"pw" form:"pw" query:"pw"`
	NAME  string `json:"Name" form:"Name" query:"Name"`
	EMAIL string `json:"Email" form:"Email" query:"Email"`
}

func SignUp(){

}
