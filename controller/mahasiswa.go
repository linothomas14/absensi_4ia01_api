package controller

import (
	"net/http"

	"github.com/linothomas14/absensi_4ia01_api/entity"
	"github.com/linothomas14/absensi_4ia01_api/helper"
	"github.com/linothomas14/absensi_4ia01_api/service"

	"github.com/gin-gonic/gin"
)

type MahasiswaController interface {
	All(context *gin.Context)
	FindByNPM(context *gin.Context)
	// Insert(context *gin.Context)
	// Update(context *gin.Context)
	// Delete(context *gin.Context)
}

type mahasiswaController struct {
	//
	mahasiswaService service.MahasiswaService
}

func NewMahasiswaController(mahasiswaServ service.MahasiswaService) MahasiswaController {
	return &mahasiswaController{
		mahasiswaService: mahasiswaServ,
	}
}

func (c *mahasiswaController) All(context *gin.Context) {

	var mahasiswas []entity.Mahasiswa = c.mahasiswaService.All()

	res := helper.BuildResponse("OK", mahasiswas)
	context.JSON(http.StatusOK, res)
}

func (c *mahasiswaController) FindByNPM(context *gin.Context) {
	npm := context.Param("npm")

	var mhs entity.Mahasiswa = c.mahasiswaService.FindByNPM(npm)
	if (mhs == entity.Mahasiswa{}) {
		res := helper.BuildErrorResponse("Data not found", "No data with given NPM", helper.EmptyObj{})
		context.JSON(http.StatusNotFound, res)
	} else {
		res := helper.BuildResponse("OK", mhs)
		context.JSON(http.StatusOK, res)
	}
}
