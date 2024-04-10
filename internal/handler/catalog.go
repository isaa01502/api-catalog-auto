package handler

import (
	"api-catalog-auto/internal/core/port/service"
	"api-catalog-auto/internal/entity"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"io"
	"net/http"
)

type Catalog struct {
	service service.ICatalog
}

func NewConsent(service service.ICatalog) *Catalog {
	return &Catalog{service: service}
}

// GetCatalog godoc
// @Summary Получить каталог
// @Description Получить каталог
// @Tags catalog
// @Accept json
// @Produce json
// @Success 200 {object} entity.Catalog{}
// @Router /api/v1/catalog [get]
// @Security Bearer
func (c *Catalog) GetCatalog(ctx *gin.Context) {
	id := ctx.Query("id")
	res, err := c.service.GetCatalog(id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, res)
	return
}

// AddCatalog godoc
// @Summary Добавить каталог
// @Description Добавить каталог
// @Tags catalog
// @Accept json
// @Produce json
// @Success 200 {object}
// @Router /api/v1/catalog [get]
// @Security Bearer
func (c *Catalog) AddCatalog(ctx *gin.Context) {
	reqBytes, err := io.ReadAll(ctx.Request.Body)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}

	var catalogRequest entity.Catalog

	if err = json.Unmarshal(reqBytes, &catalogRequest); err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}

	err = c.service.AddCatalog(catalogRequest)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, nil)
	return
}

func (c *Catalog) UpdateCatalog(ctx *gin.Context) {
	id := ctx.Query("id")

	err := c.service.UpdateCatalog(id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, nil)
	return
}

func (c *Catalog) DeleteCatalog(ctx *gin.Context) {

}
