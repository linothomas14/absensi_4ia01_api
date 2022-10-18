package repository

import (
	"log"

	"github.com/linothomas14/absensi_4ia01_api/entity"

	"gorm.io/gorm"
)

type MahasiswaRepository interface {
	InsertMahasiswa(mhs entity.Mahasiswa) entity.Mahasiswa
	AllMahasiswa() []entity.Mahasiswa
	FindByNPM(npm string) entity.Mahasiswa
}

type mahasiswaConnection struct {
	connection *gorm.DB
}

func NewMahasiswaRepository(db *gorm.DB) MahasiswaRepository {
	return &mahasiswaConnection{
		connection: db,
	}
}

func (db *mahasiswaConnection) AllMahasiswa() []entity.Mahasiswa {
	var mhs []entity.Mahasiswa
	db.connection.Find(&mhs)
	return mhs

}

func (db *mahasiswaConnection) FindByNPM(npm string) entity.Mahasiswa {
	var mhs entity.Mahasiswa

	db.connection.Where("NPM = ?", npm).Take(&mhs)
	log.Println(mhs)
	return mhs

}

func (db *mahasiswaConnection) InsertMahasiswa(m entity.Mahasiswa) entity.Mahasiswa {
	db.connection.Save(&m)
	return m
}
