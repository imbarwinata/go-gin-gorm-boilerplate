package forms

type AuthValidation struct {
	Email 		string	`json:"email" xml:"email" binding:"required"`
  Password	string	`json:"password" xml:"password" binding:"required"`
}

type AuthCheckTokenValidation struct {
	Token string `json:"token" xml:"token" binding:"required"`
}
