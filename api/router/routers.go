package api

import (
	middleware "github.com/deepmap/oapi-codegen/pkg/gin-middleware"
	"github.com/gin-gonic/gin"
	"go-backend-template/api"
	controller "go-backend-template/api/controller"
	"go-backend-template/config"
	"go-backend-template/pkg/logging"
)

// SetupHandlers sets up the router
func SetupHandlers(handlers *controller.HandlerImpl) *gin.Engine {
	swagger, err := api.GetSwagger()
	if err != nil {
		logging.Log.Error("cannot get swagger: %v", err)
		panic(err)
	}

	gin.SetMode(config.Cfg.Server.RunMode)
	router := gin.Default()
	router.Use(middleware.OapiRequestValidator(swagger))
	router = api.RegisterHandlers(router, handlers)

	return router
}
