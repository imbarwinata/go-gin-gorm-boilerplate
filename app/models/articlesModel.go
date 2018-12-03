package models

import (
	"github.com/imbarwinata/go-gin-gorm-bolerplate/app/forms"
)

//Article data struct
type Article struct {
  ID uint `json:"id"`
  Title string `json:"title"`
  Subtitle string `json:"subtitle"`
  Description string `json:"description"`
  UserID  uint `json:"userID"`
}

func (h Article) GetArticlesAll() ([]Article, error) {
	Init()
	db := GetDB()
  db.LogMode(true)
	var article []Article

  if err := db.Find(&article).Error; err != nil {
		return nil, err
  } else {
		return article, nil
	}
}

func (h Article) GetArticles(userID string) ([]Article, error) {
	Init()
	db := GetDB()
  db.LogMode(true)
	var article []Article

  if err := db.Where("user_id = ?", userID).Find(&article).Error; err != nil {
		return nil, err
  } else {
		return article, nil
	}
}

func (h Article) GetArticle(userID, id string) ([]Article, error) {
	Init()
	db := GetDB()
  db.LogMode(true)
	var user []Article

  if err := db.Where("user_id = ? AND id = ?", userID, id).First(&user, id).Error; err != nil {
		return nil, err
  } else {
		return user, nil
	}
}

func (h Article) InsertArticle(u forms.ArticleValidation) (interface{}, error) {
	Init()
	db := GetDB()
	var article = struct {
		Title string `json:"title"`
		Subtitle string `json:"subtitle"`
		Description string `json:"description"`
		UserID uint `json:"userID"`
	}{ u.Title, u.Subtitle, u.Description, u.UserID }
	// Proccess Insert
	if err := db.Table("articles").Create(&article).Error; err != nil {
		return nil, err
  } else {
		return article, nil
	}
}

func (h Article) UpdateArticle(id string, u forms.ArticleValidation) (interface{}, error) {
	Init()
	db := GetDB()
	var article Article

	if err := db.Find(&article, id).Error; err != nil {
		return nil, err
  } else {
		article.Title = u.Title
		article.Subtitle = u.Subtitle
		// Proccess Update
		if err := db.Save(&article).Error; err != nil {
				return nil, err
		} else {
				return article, nil
		}
	}
}

func (h Article) DeleteArticle(id string) (interface{}, error) {
	Init()
	db := GetDB()
  db.LogMode(true)
	var article Article

	if err := db.Find(&article, id).Error; err != nil {
		return nil, err
  }
	// Proccess Delete
	if err := db.Where("id = ?", id).Delete(&article).Error; err != nil {
			return nil, err
	} else {
			return article, nil
	}
}
