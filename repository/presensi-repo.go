package repository

import (
	"github.com/linothomas14/absensi_4ia01_api/dto"
	"github.com/linothomas14/absensi_4ia01_api/entity"

	"gorm.io/gorm"
)

type PresensiRepository interface {
	InsertPresensi(mhs entity.Presensi) (entity.Presensi, error)
	FindPresensiByMatkulAndMinggu(matkul string, minggu uint8) ([]dto.PresensiResultDTO, error)
	FindPresensi(npm string) (*entity.Presensi, error)
}

type presensiConnection struct {
	connection *gorm.DB
}

func NewPresensiRepository(db *gorm.DB) PresensiRepository {
	return &presensiConnection{
		connection: db,
	}
}
func (db *presensiConnection) FindPresensi(npm string) (*entity.Presensi, error) {
	var res *entity.Presensi

	err := db.connection.Where("NPM = ?", npm).Take(&res).Error

	if err != nil {
		return nil, err
	}

	return res, err
}
func (db *presensiConnection) FindPresensiByMatkulAndMinggu(matkul string, minggu uint8) ([]dto.PresensiResultDTO, error) {

	var mahasiswaPresensi []dto.PresensiResultDTO
	err := db.connection.Table("presensi").Select("presensi.npm, mahasiswa.nama").Joins("JOIN mahasiswa on presensi.npm = mahasiswa.npm").Where("matkul = ? AND minggu = ?", matkul, minggu).Find(&mahasiswaPresensi).Error
	if err != nil {
		return nil, err
	}

	return mahasiswaPresensi, err

}

func (db *presensiConnection) InsertPresensi(p entity.Presensi) (entity.Presensi, error) {

	err := db.connection.Save(&p).Error
	if err != nil {
		return p, err
	}
	err = db.connection.Find(&p).Error
	return p, err
}
