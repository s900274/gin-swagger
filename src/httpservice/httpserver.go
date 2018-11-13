package httpservice

import (
	"github.com/gin-gonic/gin"
	"github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
	_ "github.com/mattn/go-sqlite3"
	_ "httpservice/docs"
				"fmt"
		"define"
	"httpservice/services"
)

type HServer struct {
}

func NewHTTPServer() *HServer {
	s := &HServer{}
	return s
}

// @title Swagger Gin Test API
// @version 1.0
// @description Gin-Test

// @host
// @BasePath /
func (s *HServer) InitHttpServer() error {
	serverAddr := fmt.Sprintf("%s:%d", define.Cfg.HttpServerIp, define.Cfg.HttpServerPort)

	router := s.Router()
	gin.SetMode(gin.DebugMode)
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	return router.Run(serverAddr)
}

func (s *HServer) Router() *gin.Engine {
	r := gin.Default()

	building := r.Group("/api/building")
	{
		ginController(building)
	}

	return r
}

func ginController(building *gin.RouterGroup) {

	c := services.BuildingController{}

	company := building.Group("/company")
	{
		company.GET("", c.GetBuildingCompany)
		company.POST("", c.AddBuildingCompany)
	}
}