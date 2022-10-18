package dto

//PresensiInsertDTO is is a entity that client use when insert a new Presensi
type PresensiInsertDTO struct {
	NPM     string `json:"npm" form:"npm" binding:"required"`
	Matkul  string `json:"matkul" form:"matkul" binding:"required"`
	Tanggal string `json:"tanggal" form:"tanggal" binding:"required"`
}

type PresensiGetDTO struct {
	Matkul  string `json:"matkul" form:"matkul" binding:"required"`
	Tanggal string `json:"tanggal" form:"tanggal" binding:"required"`
}
