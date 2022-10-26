package service

import (
	"log"

	"github.com/linothomas14/absensi_4ia01_api/dto"
	"github.com/linothomas14/absensi_4ia01_api/entity"
	"github.com/linothomas14/absensi_4ia01_api/repository"

	"github.com/mashingan/smapping"
)

type PresensiService interface {
	FindByMatkulAndDate(matkul, waktu string) entity.Presensi
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

func (service *presensiService) FindByMatkulAndDate(matkul, waktu string) entity.Presensi {
	return service.presensiRepository.FindByMatkulAndDate(matkul, waktu)
}

func (service *presensiService) Insert(p dto.PresensiInsertDTO) (entity.Presensi, error) {
	presensi := entity.Presensi{}

	err := smapping.FillStruct(&presensi, smapping.MapFields(&p))

	if err != nil {
		log.Fatalf("Failed map %v: ", err)
		return entity.Presensi{}, err
	}
	res, err := service.presensiRepository.InsertPresensi(presensi)

	if err != nil {
		return entity.Presensi{}, err
	}
	return res, err
}
