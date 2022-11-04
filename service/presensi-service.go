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
	// BackupFindByMatkulAndMinggu(matkul string, minggu int) (entity.Presensi, error)
	FindByMatkulAndMinggu(matkul string, minggu uint8) (responseGetPresensi, error)
	Insert(p dto.PresensiInsertDTO) (entity.Presensi, error)
}

type presensiService struct {
	presensiRepository repository.PresensiRepository
}

func NewPresensiService(mhsRep repository.PresensiRepository) PresensiService {
	return &presensiService{
		presensiRepository: mhsRep,
	}
}

func (service *presensiService) FindByMatkulAndMinggu(matkul string, minggu uint8) (responseGetPresensi, error) {

	val, err := service.presensiRepository.FindPresensiByMatkulAndMinggu(matkul, minggu)
	res := parseGetPresensi(matkul, minggu, val)
	return res, err
}

func (service *presensiService) Insert(p dto.PresensiInsertDTO) (entity.Presensi, error) {
	presensi := entity.Presensi{}

	err := smapping.FillStruct(&presensi, smapping.MapFields(&p))

	if err != nil {
		log.Fatalf("Failed map %v: ", err)
		return entity.Presensi{}, err
	}

	isExist, err := service.presensiRepository.FindPresensi(presensi.NPM)

	if isExist != nil {
		err := errors.New(isExist.NPM + " Already present")
		return entity.Presensi{}, err
	}

	res, err := service.presensiRepository.InsertPresensi(presensi)

	if err != nil {
		return entity.Presensi{}, err
	}
	return res, err
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
