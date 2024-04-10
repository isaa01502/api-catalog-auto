package service

import (
	"api-catalog-auto/internal/entity"
)

type ICatalog interface {
	GetCatalog(id string) (catalog *entity.Catalog, err error)
	AddCatalog(request entity.Catalog) error
	UpdateCatalog(id string) error
	DeleteCatalog(id string) error
}
