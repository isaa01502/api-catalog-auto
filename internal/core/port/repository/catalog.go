package repository

import "api-catalog-auto/internal/core/dto"

type Catalog interface {
	GetCatalogByID(id string) (consent dto.CatalogDTO, err error)
	AddCatalog(catalogDTO dto.CatalogDTO) (err error)
	UpdateCatalog(catalogDTO dto.CatalogDTO) (err error)
	DeleteCatalog(id string) (err error)
}
