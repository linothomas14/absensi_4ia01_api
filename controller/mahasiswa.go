package controller

import (
	"net/http"

	"github.com/linothomas14/absensi_4ia01_api/dto"
	"github.com/linothomas14/absensi_4ia01_api/helper"
	"github.com/linothomas14/absensi_4ia01_api/model"
	"github.com/linothomas14/absensi_4ia01_api/service"

	"github.com/gin-gonic/gin"
)

type MahasiswaController interface {
	All(context *gin.Context)
	FindByNPM(context *gin.Context)
	Insert(context *gin.Context)
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

	var mahasiswas []model.Mahasiswa = c.mahasiswaService.All()
	// fmt.Println(mahasiswas)
	res := helper.BuildResponse(true, "OK", mahasiswas)
	context.JSON(http.StatusOK, res)
}

func (c *mahasiswaController) GetMahasiswa(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{
		"msg": "Yulyano Thomas Djaya",
	})
}

func (c *mahasiswaController) FindByNPM(context *gin.Context) {
	npm := context.Param("npm")

	var mhs model.Mahasiswa = c.mahasiswaService.FindByNPM(npm)
	if (mhs == model.Mahasiswa{}) {
		res := helper.BuildErrorResponse("Data not found", "No data with given NPM", helper.EmptyObj{})
		context.JSON(http.StatusNotFound, res)
	} else {
		res := helper.BuildResponse(true, "OK", mhs)
		context.JSON(http.StatusOK, res)
	}
}

func (c *mahasiswaController) Insert(context *gin.Context) {

	var mhsCreateDTO dto.MahasiswaCreateDTO
	errDTO := context.ShouldBind(&mhsCreateDTO)
	if errDTO != nil {
		res := helper.BuildErrorResponse("Failed to process request", errDTO.Error(), helper.EmptyObj{})
		context.JSON(http.StatusBadRequest, res)
	} else {
		result := c.mahasiswaService.Insert(mhsCreateDTO)
		response := helper.BuildResponse(true, "OK", result)
		context.JSON(http.StatusCreated, response)
	}

}

// 	var mahasiswa model.Mahasiswa = c.mahasiswaService.FindByNPM(npm)
// 	if (mahasiswa == model.Mahasiswa{}) {
// 		res := helper.BuildErrorResponse("Data not found", "No data with given id", helper.EmptyObj{})
// 		context.JSON(http.StatusNotFound, res)
// 	} else {
// 		res := helper.BuildResponse(true, "OK", mahasiswa)
// 		context.JSON(http.StatusOK, res)
// 	}
// }

// func (c *mahasiswaController) GetMahasiswaById(ctx *gin.Context) {
// 	ctx.JSON(http.StatusOK, gin.H{
// 		"msg": "Yulyano Thomas Djaya 123",
// 	})
// }
