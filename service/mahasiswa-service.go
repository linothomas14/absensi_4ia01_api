package service

import (
	"log"

	"github.com/linothomas14/absensi_4ia01_api/dto"
	"github.com/linothomas14/absensi_4ia01_api/model"
	"github.com/linothomas14/absensi_4ia01_api/repository"

	"github.com/mashingan/smapping"
)

type MahasiswaService interface {
	FindByNPM(npm string) model.Mahasiswa
	All() []model.Mahasiswa
	Insert(m dto.MahasiswaCreateDTO) model.Mahasiswa
}

type mahasiswaService struct {
	mahasiswaRepository repository.MahasiswaRepository
}

func NewMahasiswaService(mhsRep repository.MahasiswaRepository) MahasiswaService {
	return &mahasiswaService{
		mahasiswaRepository: mhsRep,
	}
}

func (service *mahasiswaService) All() []model.Mahasiswa {

	return service.mahasiswaRepository.AllMahasiswa()
}

func (service *mahasiswaService) FindByNPM(npm string) model.Mahasiswa {
	return service.mahasiswaRepository.FindByNPM(npm)
}

func (service *mahasiswaService) Insert(m dto.MahasiswaCreateDTO) model.Mahasiswa {
	mhs := model.Mahasiswa{}
	err := smapping.FillStruct(&mhs, smapping.MapFields(&m))
	if err != nil {
		log.Fatalf("Failed map %v: ", err)
	}
	res := service.mahasiswaRepository.InsertMahasiswa(mhs)
	return res
}
