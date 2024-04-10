package http

import (
	"api-catalog-auto/config"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/swag/example/basic/docs"
)

func (s *server) addSwaggerSettings(cfg *config.SwaggerUIConfig) {
	// Настройки swagger
	docs.SwaggerInfo.Host = cfg.Host
	docs.SwaggerInfo.Schemes = []string{"https"}
	docs.SwaggerInfo.Description = cfg.Description
	docs.SwaggerInfo.Title = cfg.PageTitle

	s.router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}
