package forms

type InsertAccountValidation struct {
  JenisKelamin  string  `json:"jenis_kelamin" xml:"jenis_kelamin" binding:"required"`
  TanggalLahir  string  `json:"tanggal_lahir" xml:"tanggal_lahir" binding:"required"`
  TempatLahir   string  `json:"tempat_lahir" xml:"tempat_lahir" binding:"required"`
  UserID        uint    `json:"userID" xml:"userID" binding:"required"`
}

type UpdateAccountValidation struct {
  JenisKelamin  string  `json:"jenis_kelamin" xml:"jenis_kelamin" binding:"required"`
  TanggalLahir  string  `json:"tanggal_lahir" xml:"tanggal_lahir" binding:"required"`
  TempatLahir   string  `json:"tempat_lahir" xml:"tempat_lahir" binding:"required"`
}
