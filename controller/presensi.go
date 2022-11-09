package controller

import (
	"net/http"

	"github.com/linothomas14/absensi_4ia01_api/dto"
	"github.com/linothomas14/absensi_4ia01_api/helper"
	"github.com/linothomas14/absensi_4ia01_api/service"

	"github.com/gin-gonic/gin"
)

type PresensiController interface {
	// All(context *gin.Context)
	FindPresensiByMatkulAndMinggu(context *gin.Context)
	Insert(context *gin.Context)
	Delete(context *gin.Context)
	// Update(context *gin.Context)
}

type presensiController struct {
	presensiService service.PresensiService
}

func NewPresensiController(presensiServ service.PresensiService) PresensiController {
	return &presensiController{
		presensiService: presensiServ,
	}
}
func (c *presensiController) Insert(context *gin.Context) {
	var presensiInsertDTO dto.PresensiDTO

	errDTO := context.ShouldBind(&presensiInsertDTO)

	if errDTO != nil {
		res := helper.BuildErrorResponse(errDTO.Error(), helper.EmptyObj{})
		context.JSON(http.StatusBadRequest, res)
		return
	}

	result, err := c.presensiService.Insert(presensiInsertDTO)

	if err != nil {
		res := helper.BuildErrorResponse(err.Error(), helper.EmptyObj{})
		context.JSON(http.StatusBadRequest, res)
		return
	}

	response := helper.BuildResponse("OK", result)
	context.JSON(http.StatusCreated, response)

}

func (c *presensiController) FindPresensiByMatkulAndMinggu(context *gin.Context) {
	var presensiGetDTO dto.PresensiGetDTO

	errDTO := context.ShouldBindJSON(&presensiGetDTO)

	if errDTO != nil {
		res := helper.BuildErrorResponse(errDTO.Error(), helper.EmptyObj{})
		context.JSON(http.StatusBadRequest, res)
		return
	}

	result, err := c.presensiService.FindByMatkulAndMinggu(presensiGetDTO.Matkul, uint8(presensiGetDTO.Minggu))

	if err != nil {
		res := helper.BuildErrorResponse(err.Error(), helper.EmptyObj{})
		context.JSON(http.StatusBadRequest, res)
		return
	}

	response := helper.BuildResponse("OK", result)
	context.JSON(http.StatusCreated, response)

}

func (c *presensiController) Delete(context *gin.Context) {
	var presensiDTO dto.PresensiDTO

	errDTO := context.ShouldBindJSON(&presensiDTO)

	if errDTO != nil {
		res := helper.BuildErrorResponse(errDTO.Error(), helper.EmptyObj{})
		context.JSON(http.StatusBadRequest, res)
		return
	}

	err := c.presensiService.Delete(presensiDTO.NPM, presensiDTO.Matkul, uint8(presensiDTO.Minggu))

	if err != nil {
		res := helper.BuildErrorResponse(err.Error(), helper.EmptyObj{})
		context.JSON(http.StatusBadRequest, res)
		return
	}

	response := helper.BuildResponse("Data berhasil dihapus", "")
	context.JSON(http.StatusCreated, response)

}
