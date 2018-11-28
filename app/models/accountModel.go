package models

//Article data struct
type Account struct {
  ID uint `json:"id"`
  JenisKelamin string `json:"jenis_kelamin"`
  TanggalLahir string `json:"tanggal_lahir"`
  TempatLahir string `json:"tempat_lahir"`
  UserID  uint
  // User models.User
}
