package repository

import (
	"github.com/linothomas14/absensi_4ia01_api/entity"

	"gorm.io/gorm"
)

type PresensiRepository interface {
	InsertPresensi(mhs entity.Presensi) (entity.Presensi, error)
	FindByMatkulAndMinggu(matkul string, minggu int) (entity.Presensi, error)
	FindPresensiByMatkulAndMinggu(matkul string, minggu int) ([]entity.Presensi, error)
}

type presensiConnection struct {
	connection *gorm.DB
}

func NewPresensiRepository(db *gorm.DB) PresensiRepository {
	return &presensiConnection{
		connection: db,
	}
}

func (db *presensiConnection) FindPresensiByMatkulAndMinggu(matkul string, minggu int) ([]entity.Presensi, error) {

	var presensis []entity.Presensi

	err := db.connection.Where("matkul = ? AND minggu = ?", matkul, minggu).Take(&presensis).Error

	return presensis, err

}

func (db *presensiConnection) FindByMatkulAndMinggu(matkul string, minggu int) (entity.Presensi, error) {
	var presensi entity.Presensi
	err := db.connection.Where("matkul = ? AND minggu = ?", matkul, minggu).Take(&presensi).Error

	return presensi, err

}

func (db *presensiConnection) InsertPresensi(p entity.Presensi) (entity.Presensi, error) {

	err := db.connection.Save(&p).Error
	if err != nil {
		return p, err
	}
	err = db.connection.Find(&p).Error
	return p, err
}
