package http

import (
	"api-catalog-auto/internal/core/port/service"
	"api-catalog-auto/internal/handler"
)

type Handlers struct {
	catalog *handler.Catalog
}

func NewHandlers(catalogSrv service.ICatalog) *Handlers {
	return &Handlers{catalog: handler.NewConsent(catalogSrv)}
}
