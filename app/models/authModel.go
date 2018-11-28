package models

func (h User) Login(email string, password string) ([]User, error) {
	Init()
	db := GetDB()
  // db.LogMode(true)
	var auth []User

  if err := db.Where(&User{
		Email:    email,
		Password: password,
	}).First(&auth).Error; err != nil {
    return nil, err
  }

	return auth, nil
}
