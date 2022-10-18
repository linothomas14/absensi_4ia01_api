package controller

import (
	"net/http"

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

	var res helper.Response
	mahasiswas, err := c.mahasiswaService.All()

	if err != nil {
		res = helper.BuildErrorResponse(err.Error(), mahasiswas)
	}

	res = helper.BuildResponse("OK", mahasiswas)
	context.JSON(http.StatusOK, res)
}

func (c *mahasiswaController) FindByNPM(context *gin.Context) {
	npm := context.Param("npm")

	mhs, err := c.mahasiswaService.FindByNPM(npm)
	if err != nil {
		res := helper.BuildErrorResponse(err.Error(), helper.EmptyObj{})
		context.JSON(http.StatusNotFound, res)
	} else {
		res := helper.BuildResponse("OK", mhs)
		context.JSON(http.StatusOK, res)
	}
}
