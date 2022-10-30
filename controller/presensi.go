package controller

import (
	"log"
	"net/http"

	"github.com/linothomas14/absensi_4ia01_api/dto"
	"github.com/linothomas14/absensi_4ia01_api/helper"
	"github.com/linothomas14/absensi_4ia01_api/service"

	"github.com/gin-gonic/gin"
)

type PresensiController interface {
	// All(context *gin.Context)
	FindByMatkulAndMinggu(context *gin.Context)
	Insert(context *gin.Context)
	// Update(context *gin.Context)
	// Delete(context *gin.Context)
}

type presensiController struct {
	//
	presensiService service.PresensiService
}

func NewPresensiController(presensiServ service.PresensiService) PresensiController {
	return &presensiController{
		presensiService: presensiServ,
	}
}
func (c *presensiController) Insert(context *gin.Context) {
	var presensiInsertDTO dto.PresensiInsertDTO
	log.Println(presensiInsertDTO)
	errDTO := context.ShouldBind(&presensiInsertDTO)
	log.Println(presensiInsertDTO)
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

func (c *presensiController) FindByMatkulAndMinggu(context *gin.Context) {
	var presensiGetDTO dto.PresensiGetDTO

	errDTO := context.ShouldBindJSON(&presensiGetDTO)

	if errDTO != nil {
		res := helper.BuildErrorResponse(errDTO.Error(), helper.EmptyObj{})
		context.JSON(http.StatusBadRequest, res)
		return
	}
	log.Println(presensiGetDTO.Matkul, presensiGetDTO.Minggu)
	result, err := c.presensiService.FindByMatkulAndMinggu(presensiGetDTO.Matkul, int(presensiGetDTO.Minggu))

	if err != nil {
		res := helper.BuildErrorResponse(err.Error(), helper.EmptyObj{})
		context.JSON(http.StatusBadRequest, res)
		return
	}

	response := helper.BuildResponse("OK", result)
	context.JSON(http.StatusCreated, response)

}
