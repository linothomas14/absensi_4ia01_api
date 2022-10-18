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
	FindByMatkulAndDate(context *gin.Context)
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

	errDTO := context.ShouldBind(&presensiInsertDTO)

	if errDTO != nil {
		res := helper.BuildErrorResponse("Failed to process request", errDTO.Error(), helper.EmptyObj{})
		context.JSON(http.StatusBadRequest, res)
	} else {

		result := c.presensiService.Insert(presensiInsertDTO)
		response := helper.BuildResponse("OK", result)
		context.JSON(http.StatusCreated, response)
	}
}

func (c *presensiController) FindByMatkulAndDate(context *gin.Context) {

}
