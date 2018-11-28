package forms

type AuthValidation struct {
	Email 		string	`json:"email" xml:"email" binding:"required"`
  Password	string	`json:"password" xml:"password" binding:"required"`
}
