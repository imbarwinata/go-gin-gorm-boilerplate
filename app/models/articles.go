package models

//Article data struct
type Article struct {
  ID uint `json:"id"`
  Title string `json:"title"`
  Subtitle string `json:"subtitle"`
  Description string `json:"description"`
  UserID  uint
  // User models.User
}
