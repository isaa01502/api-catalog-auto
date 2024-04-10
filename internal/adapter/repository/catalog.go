package repository

import (
	"api-catalog-auto/internal/core/dto"
	"api-catalog-auto/internal/core/port/repository"
	"api-catalog-auto/internal/entity"
	"errors"
	"gorm.io/gorm"
)

type catalog struct {
	db *gorm.DB
}

func NewTemplate(dbIn repository.IDB) (repository.Catalog, error) {

	dbImp, ok := dbIn.(*database)
	if !ok {
		return nil, errors.New("mock error")
	}

	return &catalog{db: dbImp.gormDB}, nil
}

func (c catalog) GetCatalogByID(id string) (catalogDTO dto.CatalogDTO, err error) {
	catalog := entity.Catalog{}
	err = c.db.Model(&entity.Catalog{}).Where("id = ?", id).Find(&catalog).Error
	if err != nil {
		return catalogDTO, err
	}

	catalogDTO.ID = catalog.ID
	catalogDTO.Mark = catalog.Mark
	catalogDTO.Owner = catalog.Owner
	catalogDTO.RegNum = catalog.RegNum

	return
}

func (c catalog) AddCatalog(catalogDTO dto.CatalogDTO) (err error) {
	catalog := entity.Catalog{
		Mark:   catalogDTO.Mark,
		Owner:  catalogDTO.Owner,
		RegNum: catalogDTO.RegNum,
	}

	err = c.db.Create(&catalog).Error
	if err != nil {
		return err
	}

	return
}

func (c catalog) UpdateCatalog(catalogDTO dto.CatalogDTO) (err error) {
	catalog := entity.Catalog{
		ID:     catalogDTO.ID,
		Mark:   catalogDTO.Mark,
		Owner:  catalogDTO.Owner,
		RegNum: catalogDTO.RegNum,
	}

	err = c.db.Model(&entity.Catalog{}).Where("id = ?", catalog.ID).Updates(&catalog).Error
	if err != nil {
		return err
	}

	return
}

func (c catalog) DeleteCatalog(id string) (err error) {
	err = c.db.Model(&entity.Catalog{}).Where("id = ?", id).Delete(&entity.Catalog{}).Error
	if err != nil {
		return err
	}
	return
}
