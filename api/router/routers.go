package api

import (
	oapimiddleware "github.com/deepmap/oapi-codegen/pkg/middleware"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"go-backend-template/api"
	controller "go-backend-template/api/controller"
	"go-backend-template/pkg/logging"
)

// SetupHandlers sets up the router
func SetupHandlers(handlers *controller.HandlerImpl) *echo.Echo {
	swagger, err := api.GetSwagger()
	if err != nil {
		logging.Log.Error("cannot get swagger: %v", err)
		panic(err)
	}

	e := echo.New()
	e.Use(middleware.CORS())
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "[ECHO] ${time_rfc3339} | ${status} | ${latency_human} | ${id} | ${remote_ip} | ${user_agent} | ${method} | ${uri} | error=${error}\n",
	}))
	e.Use(middleware.RequestID())
	e.Use(oapimiddleware.OapiRequestValidator(swagger))
	api.RegisterHandlers(e, handlers)

	return e
}
