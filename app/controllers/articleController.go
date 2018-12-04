package controllers

import (
  "github.com/gin-gonic/gin"
	"github.com/imbarwinata/go-gin-gorm-boilerplate/app/forms"
	"github.com/imbarwinata/go-gin-gorm-boilerplate/app/models"
)
type ArticleController struct{}
type Article struct{
  title string
  subtitle string
  description string
  userid uint
}
var articleModel = new(models.Article)

func (u ArticleController) GetsAll(c *gin.Context) {
  status := 500
  api_key := c.Request.Header.Get("Api-Key")
	api_key_match := getKey()
	articles, err := articleModel.GetArticlesAll()

	if err != nil || api_key != api_key_match {
		if api_key != api_key_match {
			status = 401
		}
		c.JSON(status, gin.H{
			"message": "Kesalahan saat mengambil data article",
			"status": status,
		})
		c.Abort()
		return
	}
	c.JSON(200, gin.H{"message": "Success", "status": 200, "articles": articles})
	return
}

func (u ArticleController) Gets(c *gin.Context) {
  status := 500
  userID := c.Param("id")
	api_key := c.Request.Header.Get("Api-Key")
	api_key_match := getKey()
	articles, err := articleModel.GetArticles(userID)

	if err != nil || api_key != api_key_match {
		if api_key != api_key_match {
			status = 401
		}
		c.JSON(status, gin.H{
			"message": "Kesalahan saat mengambil data article",
			"status": status,
		})
		c.Abort()
		return
	}
	c.JSON(200, gin.H{"message": "Success", "status": 200, "articles": articles})
	return
}

func (u ArticleController) Get(c *gin.Context) {
  status := 200
	userID := c.Param("id")
  articleID := c.Param("articleid")
	api_key := c.Request.Header.Get("Api-Key")
	api_key_match := getKey()
	articles, err := articleModel.GetArticle(userID, articleID)

	if err != nil || len(articles) < 1 || api_key != api_key_match {
			if api_key != api_key_match {
				status = 401
			}
			c.JSON(status, gin.H{
				"message": "Kesalahan saat mengambil data article",
				"status": status,
			})
			c.Abort()
			return
	}
	c.JSON(200, gin.H{"message": "Success", "status": 200, "article": articles[0]})
	return
}

func (u ArticleController) Insert(c *gin.Context) {
	defer catch()
	api_key := c.Request.Header.Get("Api-Key")
	api_key_match := getKey()
	var validArticle forms.ArticleValidation

	if err = c.BindJSON(&validArticle); err != nil || api_key != api_key_match {
		if api_key != api_key_match {
			panic("dibutuhkan api key yang benar untuk mengakses ini")
		}
		c.JSON(500, gin.H{ "error":  err.Error(), "status": 500 })
		return
	}
	article, err := articleModel.InsertArticle(validArticle)
  if err != nil {
    c.JSON(200, gin.H{ "error": "Tidak dapat menambahkan data article", "status": 200 })
    return
  }
	c.JSON(200, gin.H{"message": "Success", "status": 200, "inserted": article})
	return
}

func (u ArticleController) Update(c *gin.Context) {
	defer catch()
	var validArticle forms.ArticleValidation
	api_key := c.Request.Header.Get("Api-Key")
	api_key_match := getKey()
	articleID := c.Param("articleid")

	if err = c.BindJSON(&validArticle); err != nil || api_key != api_key_match {
			if api_key != api_key_match {
				panic("dibutuhkan api key yang benar untuk mengakses ini")
			}
			c.JSON(500, gin.H{ "error":  err.Error(), "status": 500 })
			return
	}
	user, err := articleModel.UpdateArticle(articleID, validArticle)
  if err != nil {
    c.JSON(200, gin.H{ "error":  "Tidak dapat memperbaharui data article", "status": 200 })
    return
  }
	c.JSON(200, gin.H{"message": "Success", "status": 200, "updated": user})
	return
}

func (u ArticleController) Delete(c *gin.Context) {
	defer catch()
	api_key := c.Request.Header.Get("Api-Key")
	api_key_match := getKey()
  articleID := c.Param("articleid")
	article, err := articleModel.DeleteArticle(articleID)

	if err != nil || api_key != api_key_match{
			if api_key != api_key_match {
				panic("dibutuhkan api key yang benar untuk mengakses ini")
			}
			c.JSON(200, gin.H{ "message": err.Error(), "status": 500 })
      return
	}
  c.JSON(200, gin.H{"message": "Success", "status": 200, "deleted": article})
	return
}
