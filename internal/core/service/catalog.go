package service

import (
	"api-catalog-auto/config"
	"api-catalog-auto/internal/core/port/repository"
	"api-catalog-auto/internal/core/port/service"
	"api-catalog-auto/internal/entity"
	"bytes"
	"encoding/json"
	"fmt"
	log "go.uber.org/zap"
	"io"
	"net/http"
)

type catalog struct {
	cfg *config.Config
	rep repository.Catalog
}

func NewCatalog(cfg *config.Config, rep repository.Catalog) service.ICatalog {
	return &catalog{
		cfg: cfg,
		rep: rep,
	}
}

func (c catalog) GetCatalog(id string) (catalog *entity.Catalog, err error) {
	client := &http.Client{}

	url := fmt.Sprintf("%v?id=%v", c.cfg.CatalogUrls.Catalog, id)
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		log.Error(err)
		return
	}

	resp, err := client.Do(req)
	if err != nil {
		log.Error(err)
		return
	}
	defer resp.Body.Close()

	reqBytes, err := io.ReadAll(resp.Body)
	if err = json.Unmarshal(reqBytes, &catalog); err != nil {
		log.Error(err)
		return
	}

	return
}

func (c catalog) AddCatalog(request entity.Catalog) error {
	client := &http.Client{}

	reqCatalog := entity.Catalog{
		RegNum: request.RegNum,
		Mark:   request.Mark,
		Model:  request.Model,
	}

	reqByte, err := json.Marshal(reqCatalog)
	if err != nil {
		log.Error(err)
		return err
	}

	url := c.cfg.CatalogUrls.Catalog
	req, err := http.NewRequest(http.MethodPost, url, bytes.NewBuffer(reqByte))
	if err != nil {
		log.Error(err)
		return err
	}

	resp, err := client.Do(req)
	if err != nil {
		log.Error(err)
		return err
	}
	defer resp.Body.Close()

	return nil
}

func (c catalog) UpdateCatalog(id string) error {
	client := &http.Client{}

	url := fmt.Sprintf("%v?id=%v", c.cfg.CatalogUrls.Catalog, id)
	req, err := http.NewRequest(http.MethodPut, url, nil)
	if err != nil {
		log.Error(err)
		return err
	}

	resp, err := client.Do(req)
	if err != nil {
		log.Error(err)
		return err
	}
	defer resp.Body.Close()

	return nil
}

func (c catalog) DeleteCatalog(id string) error {
	client := &http.Client{}

	url := fmt.Sprintf("%v?id=%v", c.cfg.CatalogUrls.Catalog, id)
	req, err := http.NewRequest(http.MethodDelete, url, nil)
	if err != nil {
		log.Error(err)
		return err
	}

	resp, err := client.Do(req)
	if err != nil {
		log.Error(err)
		return err
	}
	defer resp.Body.Close()

	return nil
}
