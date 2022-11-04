package repository

import (
	"github.com/linothomas14/absensi_4ia01_api/entity"

	"gorm.io/gorm"
)

type MahasiswaRepository interface {
	InsertMahasiswa(mhs entity.Mahasiswa) entity.Mahasiswa
	AllMahasiswa() ([]entity.Mahasiswa, error)
	FindByNPM(npm string) (entity.Mahasiswa, error)
}

type mahasiswaConnection struct {
	connection *gorm.DB
}

func NewMahasiswaRepository(db *gorm.DB) MahasiswaRepository {
	return &mahasiswaConnection{
		connection: db,
	}
}

func (db *mahasiswaConnection) AllMahasiswa() ([]entity.Mahasiswa, error) {
	var mhs []entity.Mahasiswa
	err := db.connection.Find(&mhs).Error
	if err != nil {
		return []entity.Mahasiswa{}, err
	}
	return mhs, err

}

func (db *mahasiswaConnection) FindByNPM(npm string) (entity.Mahasiswa, error) {
	var mhs entity.Mahasiswa

	err := db.connection.Where("NPM = ?", npm).Take(&mhs).Error
	if err != nil {
		return mhs, err
	}
	return mhs, err

}

func (db *mahasiswaConnection) InsertMahasiswa(m entity.Mahasiswa) entity.Mahasiswa {
	db.connection.Save(&m)
	return m
}
