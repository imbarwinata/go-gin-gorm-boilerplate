package forms

type SiswaValidation struct {
  NomorInduk    string  `json:"nomorInduk" xml:"nomorInduk" binding:"required"`
  NISN          string  `json:"nisn" xml:"nisn" binding:"required"`
  Nama          string  `json:"nama" xml:"nama" binding:"required"`
  TanggalLahir  string  `json:"tanggalLahir" xml:"tanggalLahir" binding:"required"`
  TempatLahir   string  `json:"tempatLahir" xml:"tempatLahir" binding:"required"`
  JenisKelamin  string  `json:"jenisKelamin" xml:"jenisKelamin" binding:"required"`
  Alamat        string  `json:"alamat" xml:"alamat" binding:"required"`
  Kelas         string  `json:"kelas" xml:"kelas" binding:"required"`
  Status        string  `json:"status" xml:"status" binding:"required"`
  OrtuNama      string  `json:"ortuNama" xml:"ortuNama"`
  OrtuHandphone string  `json:"ortuHandphone" xml:"ortuHandphone"`
}
