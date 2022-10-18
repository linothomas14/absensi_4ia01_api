package main

import (
	"github.com/linothomas14/absensi_4ia01_api/config"
	"github.com/linothomas14/absensi_4ia01_api/controller"
	"github.com/linothomas14/absensi_4ia01_api/repository"
	"github.com/linothomas14/absensi_4ia01_api/service"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var (
	db                  *gorm.DB                       = config.SetupDatabaseConnection()
	mahasiswaRepository repository.MahasiswaRepository = repository.NewMahasiswaRepository(db)
	mahasiswaService    service.MahasiswaService       = service.NewMahasiswaService(mahasiswaRepository)
	mahasiswaController controller.MahasiswaController = controller.NewMahasiswaController(mahasiswaService)
)

func PingHandler(c *gin.Context) {
	c.JSON(200, gin.H{
		"msg": "pong",
	})
}

func main() {
	// config.MigrateDB(db)
	defer config.CloseDatabaseConnection(db)
	r := gin.Default()

	mhsRoutes := r.Group("api/mahasiswa")
	{
		mhsRoutes.GET("/", mahasiswaController.All)
		mhsRoutes.POST("/", mahasiswaController.Insert)
		mhsRoutes.GET("/:npm", mahasiswaController.FindByNPM)
	}
	r.GET("api/ping", mahasiswaController.All)
	r.Run()
}
