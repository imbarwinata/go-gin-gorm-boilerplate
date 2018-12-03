package forms

type ArticleValidation struct {
  Title         string  `json:"title" xml:"title" binding:"required"`
  Subtitle      string  `json:"subtitle" xml:"subtitle" binding:"required"`
  Description   string  `json:"description" xml:"description" binding:"required"`
  UserID        uint    `json:"userID" xml:"userID" binding:"required"`
}
