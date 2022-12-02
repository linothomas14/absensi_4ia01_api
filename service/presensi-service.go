package service

import (
	"errors"
	"log"

	"github.com/linothomas14/absensi_4ia01_api/dto"
	"github.com/linothomas14/absensi_4ia01_api/entity"
	"github.com/linothomas14/absensi_4ia01_api/repository"

	"github.com/mashingan/smapping"
)

type PresensiService interface {
	FindByMatkulAndMinggu(matkul string, minggu uint8) (responseGetPresensi, error)
	Insert(p dto.PresensiDTO) (entity.Presensi, error)
	Delete(npm string, matkul string, minggu uint8) error
}

type presensiService struct {
	presensiRepository  repository.PresensiRepository
	mahasiswaRepository repository.MahasiswaRepository
}

func NewPresensiService(preRep repository.PresensiRepository, mhsRep repository.MahasiswaRepository) PresensiService {
	return &presensiService{
		presensiRepository:  preRep,
		mahasiswaRepository: mhsRep,
	}
}

func (service *presensiService) FindByMatkulAndMinggu(matkul string, minggu uint8) (responseGetPresensi, error) {

	val, err := service.presensiRepository.FindPresensiByMatkulAndMinggu(matkul, minggu)
	res := parseGetPresensi(matkul, minggu, val)
	return res, err
}

func (service *presensiService) Insert(p dto.PresensiDTO) (entity.Presensi, error) {
	presensi := entity.Presensi{}

	err := smapping.FillStruct(&presensi, smapping.MapFields(&p))

	if err != nil {
		log.Fatalf("Failed map %v: ", err)
		return entity.Presensi{}, err
	}
	// Check if mahasiswa not in 4IA01
	mhs, err := service.mahasiswaRepository.FindByNPM(presensi.NPM)

	if mhs.NPM == "" {
		err := errors.New(mhs.NPM + " not register in 4IA01")
		return entity.Presensi{}, err
	}

	mhs2, err := service.presensiRepository.FindPresensi(presensi.NPM, presensi.Matkul, presensi.Minggu)

	if mhs2 != nil {
		err := errors.New(mhs2.NPM + " Already present")
		return entity.Presensi{}, err
	}

	res, err := service.presensiRepository.InsertPresensi(presensi)

	if err != nil {
		return entity.Presensi{}, err
	}
	return res, err
}

// MASIH BUG
func (service *presensiService) Delete(npm string, matkul string, minggu uint8) error {

	res, err := service.presensiRepository.FindPresensi(npm, matkul, minggu)

	if res == nil {
		return err
	}
	err = service.presensiRepository.DeletePresensi(npm, matkul, minggu)
	return err
}

func parseGetPresensi(matkul string, minggu uint8, mahasiswa []dto.PresensiResultDTO) responseGetPresensi {

	var res responseGetPresensi

	res.Matkul = matkul
	res.Minggu = minggu
	res.Mahasiswa = mahasiswa

	return res
}

type responseGetPresensi struct {
	Matkul    string                  `json:"matkul"`
	Minggu    uint8                   `json:"minggu"`
	Mahasiswa []dto.PresensiResultDTO `json:"mahasiswa"`
}
