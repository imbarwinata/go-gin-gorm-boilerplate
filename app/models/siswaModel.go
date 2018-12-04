package models

import (
	"github.com/imbarwinata/go-gin-gorm-boilerplate/app/forms"
)

//Siswa data struct
type Siswa struct {
  ID uint `json:"id"`
  NomorInduk string `json:"nomorInduk"`
  NISN string `json:"nisn"`
  Nama  string `json:"nama"`
  TanggalLahir  string `json:"tanggalLahir"`
  TempatLahir  string `json:"tempatLahir"`
  JenisKelamin  string `json:"jenisKelamin"`
  Alamat  string `json:"alamat"`
  Kelas  string `json:"kelas"`
  Status  string `json:"status"`
  OrtuNama  string `json:"ortuNama"`
  OrtuHandphone string `json:"ortuHandphone"`
}

func (h Siswa) GetSiswasAll() ([]Siswa, error) {
	Init()
	db := GetDB()
	var siswa []Siswa

  if err := db.Find(&siswa).Error; err != nil {
		return nil, err
  } else {
		return siswa, nil
	}
}

func (h Siswa) GetSiswas() ([]Siswa, error) {
	Init()
	db := GetDB()
	var siswa []Siswa

  if err := db.Find(&siswa).Error; err != nil {
		return nil, err
  } else {
		return siswa, nil
	}
}

func (h Siswa) GetSiswa(id string) ([]Siswa, error, bool) {
	Init()
	db := GetDB()
	var siswa []Siswa

  if err := db.First(&siswa, id).Error; err != nil {
		return nil, err, false
  }
  if len(siswa) < 1 {
    return nil, err, false
  }
	return siswa, nil, true
}

func (h Siswa) InsertSiswa(u forms.SiswaValidation) (interface{}, error) {
	Init()
	db := GetDB()
	var siswa = struct {
    NomorInduk    string  `json:"nomorInduk"`
    NISN          string  `json:"nisn"`
    Nama          string  `json:"nama"`
    TanggalLahir  string  `json:"tanggalLahir"`
    TempatLahir   string  `json:"tempatLahir"`
    JenisKelamin  string  `json:"jenisKelamin"`
    Alamat        string  `json:"alamat"`
    Kelas         string  `json:"kelas"`
    Status        string  `json:"status"`
    OrtuNama      string  `json:"ortuNama"`
    OrtuHandphone string  `json:"ortuHandphone"`
	}{ u.NomorInduk, u.NISN, u.Nama, u.TanggalLahir, u.TempatLahir, u.JenisKelamin, u.Alamat, u.Kelas, u.Status, u.OrtuNama, u.OrtuHandphone }
	// Proccess Insert
	if err := db.Table("siswas").Create(&siswa).Error; err != nil {
		return nil, err
  } else {
		return siswa, nil
	}
}

func (h Siswa) UpdateSiswa(id string, u forms.SiswaValidation) (interface{}, error) {
	Init()
	db := GetDB()
	var siswa Siswa

	if err := db.Find(&siswa, id).Error; err != nil {
		return nil, err
  } else {
		siswa.NomorInduk    = u.NomorInduk
    siswa.NISN          = u.NISN
    siswa.Nama          = u.Nama
    siswa.TanggalLahir  = u.TanggalLahir
    siswa.TempatLahir   = u.TempatLahir
    siswa.JenisKelamin  = u.JenisKelamin
    siswa.Alamat        = u.Alamat
    siswa.Kelas         = u.Kelas
    siswa.Status        = u.Status
    siswa.OrtuNama      = u.OrtuNama
		siswa.OrtuHandphone = u.OrtuHandphone
		// Proccess Update
		if err := db.Save(&siswa).Error; err != nil {
				return nil, err
		} else {
				return siswa, nil
		}
	}
}

func (h Siswa) DeleteSiswa(id string) (interface{}, error) {
	Init()
	db := GetDB()
	var siswa Siswa

	if err := db.Find(&siswa, id).Error; err != nil {
		return nil, err
  }
	// Proccess Delete
	if err := db.Where("id = ?", id).Delete(&siswa).Error; err != nil {
			return nil, err
	} else {
			return siswa, nil
	}
}
