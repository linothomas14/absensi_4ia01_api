package repository

import (
	"github.com/linothomas14/absensi_4ia01_api/entity"

	"gorm.io/gorm"
)

type PresensiRepository interface {
	InsertPresensi(mhs entity.Presensi) (entity.Presensi, error)
	FindByMatkulAndDate(matkul, waktu string) entity.Presensi
}

type presensiConnection struct {
	connection *gorm.DB
}

func NewPresensiRepository(db *gorm.DB) PresensiRepository {
	return &presensiConnection{
		connection: db,
	}
}

func (db *presensiConnection) FindByMatkulAndDate(matkul, waktu string) entity.Presensi {
	var presensi entity.Presensi

	db.connection.Where("matkul = ?", matkul).Take(&presensi)

	return presensi

}

func (db *presensiConnection) InsertPresensi(p entity.Presensi) (entity.Presensi, error) {

	err := db.connection.Save(&p).Error
	err = db.connection.Preload("Mahasiswa").Find(&p).Error
	return p, err
}
