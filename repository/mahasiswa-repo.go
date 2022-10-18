package repository

import (
	"log"

	"github.com/linothomas14/absensi_4ia01_api/model"

	"gorm.io/gorm"
)

type MahasiswaRepository interface {
	InsertMahasiswa(mhs model.Mahasiswa) model.Mahasiswa
	AllMahasiswa() []model.Mahasiswa
	FindByNPM(npm string) model.Mahasiswa
}

type mahasiswaConnection struct {
	connection *gorm.DB
}

func NewMahasiswaRepository(db *gorm.DB) MahasiswaRepository {
	return &mahasiswaConnection{
		connection: db,
	}
}

func (db *mahasiswaConnection) AllMahasiswa() []model.Mahasiswa {
	var mhs []model.Mahasiswa
	db.connection.Find(&mhs)
	return mhs

}

func (db *mahasiswaConnection) FindByNPM(npm string) model.Mahasiswa {
	var mhs model.Mahasiswa

	db.connection.Where("NPM = ?", npm).Take(&mhs)
	log.Println(mhs)
	return mhs

}

func (db *mahasiswaConnection) InsertMahasiswa(m model.Mahasiswa) model.Mahasiswa {
	db.connection.Save(&m)
	return m
}
