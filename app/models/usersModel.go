package models

import (
	"github.com/imbarwinata/go-gin-gorm-bolerplate/app/forms"
)

//User data struct
type User struct {
	ID 				uint 		`json:"id"`
  Email 		string 	`json:"email"`
  Password 	string 	`json:"-"`
  FirstName string 	`json:"firstname"`
  LastName 	string 	`json:"lastname"`
	Articles 	[]Article
	Account Account
}

func (h User) GetUsers() ([]User, error) {
	Init()
	db := GetDB()
	var user []User

  if err := db.Find(&user).Error; err != nil {
		return nil, err
  } else {
		// Associate : Has Many (Articles) && Has One (Account)
		for i, _ := range user {
	    db.Model(user[i]).Related(&user[i].Articles).Related(&user[i].Account)
		}
		return user, nil
	}
}

func (h User) GetUser(id string) ([]User, error) {
	Init()
	db := GetDB()
	var user []User

  if err := db.First(&user, id).Error; err != nil {
		return nil, err
  } else {
		// Associate : Has Many (Articles) && Has One (Account)
		for i, _ := range user {
	    db.Model(user[i]).Related(&user[i].Articles).Related(&user[i].Account)
		}
		return user, nil
	}
}

func (h User) InsertUser(u forms.AddUserValidation) (interface{}, error) {
	Init()
	db := GetDB()
	var user = struct {
		Email string `json:"email"`
		FirstName string `json:"firstname"`
		LastName string `json:"lastname"`
		Password string `json:"password"`
	}{ u.Email, u.FirstName, u.LastName, u.Password }
	// Proccess Insert
	if err := db.Table("users").Create(&user).Error; err != nil {
		return nil, err
  } else {
		return user, nil
	}
}

func (h User) UpdateUser(id string, u forms.UpdateUserValidation) (interface{}, error) {
	Init()
	db := GetDB()
	var user User

	if err := db.Find(&user, id).Error; err != nil {
		return nil, err
  } else {
		user.FirstName = u.FirstName
		user.LastName = u.LastName
		// Proccess Update
		if err := db.Save(&user).Error; err != nil {
				return nil, err
		} else {
				return user, nil
		}
	}
}

func (h User) DeleteUser(id string) (interface{}, error) {
	Init()
	db := GetDB()
	var user User

	if err := db.Find(&user, id).Error; err != nil {
		return nil, err
  } else {
		// Proccess Delete
		if err := db.Where("id = ?", id).Delete(&user).Error; err != nil {
				return nil, err
		} else {
				return user, nil
		}
	}
}
