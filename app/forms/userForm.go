package forms

type AddUserValidation struct {
	Email 			string	 `json:"email" xml:"email" binding:"required"`
	Password 		string	 `json:"password" xml:"password" binding:"required"`
	FirstName 	string	 `json:"firstname" xml:"firstname" binding:"required"`
  LastName 		string	 `json:"lastname" xml:"lastname" binding:"required"`
}

type UpdateUserValidation struct {
	FirstName 	string	 `json:"firstname" xml:"firstname" binding:"required"`
  LastName 		string	 `json:"lastname" xml:"lastname" binding:"required"`
}
