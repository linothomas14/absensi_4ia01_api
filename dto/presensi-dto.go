package dto

//PresensiInsertDTO is is a entity that client use when insert a new Presensi
type PresensiDTO struct {
	NPM    string `json:"npm" form:"npm" binding:"required"`
	Matkul string `json:"matkul" form:"matkul" binding:"required"`
	Minggu uint8  `json:"minggu" form:"minggu" binding:"required"`
}

type PresensiGetDTO struct {
	Matkul string `json:"matkul" form:"matkul" binding:"required"`
	Minggu uint8  `json:"minggu" form:"minggu" binding:"required"`
}

type PresensiResultDTO struct {
	Nama string `json:"nama"`
	NPM  string `json:"npm"`
}
