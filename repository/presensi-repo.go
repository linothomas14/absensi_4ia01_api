package repository

import (
	"github.com/linothomas14/absensi_4ia01_api/dto"
	"github.com/linothomas14/absensi_4ia01_api/entity"

	"gorm.io/gorm"
)

type PresensiRepository interface {
	InsertPresensi(mhs entity.Presensi) (entity.Presensi, error)
	FindPresensiByMatkulAndMinggu(matkul string, minggu uint8) ([]dto.PresensiResultDTO, error)
	FindPresensi(npm string, matkul string, minggu uint8) (*entity.Presensi, error)
	DeletePresensi(npm string, matkul string, minggu uint8) error
}

type presensiConnection struct {
	connection *gorm.DB
}

func NewPresensiRepository(db *gorm.DB) PresensiRepository {
	return &presensiConnection{
		connection: db,
	}
}
func (db *presensiConnection) FindPresensi(npm string, matkul string, minggu uint8) (*entity.Presensi, error) {
	var res *entity.Presensi

	err := db.connection.Where("NPM = ? AND matkul = ? AND minggu = ?", npm, matkul, minggu).Take(&res).Error

	if err != nil {
		return nil, err
	}

	return res, err
}
func (db *presensiConnection) FindPresensiByMatkulAndMinggu(matkul string, minggu uint8) ([]dto.PresensiResultDTO, error) {

	var mahasiswaPresensi []dto.PresensiResultDTO
	err := db.connection.Table("presensi").Select("presensi.npm, mahasiswa.nama").Joins("JOIN mahasiswa on presensi.npm = mahasiswa.npm").Where("matkul = ? AND minggu = ?", matkul, minggu).Order("nama asc").Find(&mahasiswaPresensi).Error
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

func (db *presensiConnection) DeletePresensi(npm string, matkul string, minggu uint8) error {
	var p *entity.Presensi
	err := db.connection.Where("matkul = ? AND minggu = ? AND npm = ?", matkul, minggu, npm).Delete(&p).Error
	if err != nil {
		return err
	}

	return err
}
