package variants

import (
	"log"
	"net/http"
	"strconv"

	"github.com/bpradana/devcamp_pe_bintangpradanaerlanggaputra/pkg/domain"
	"github.com/labstack/echo/v4"
)

type VariantHandler struct {
	usecase domain.VariantUsecase
}

func NewHandler(e *echo.Group, usecase domain.VariantUsecase) {
	h := &VariantHandler{
		usecase: usecase,
	}

	// Routes
	e.GET("/variants", h.GetAll)
	e.GET("/variants/:id", h.GetByID)
	e.POST("/variants", h.Create)
	e.PUT("/variants/:id", h.Update)
	e.DELETE("/variants/:id", h.Delete)
}

func (h *VariantHandler) GetAll(c echo.Context) error {
	products, err := h.usecase.GetAll()
	if err != nil {
		log.Println("[VariantHandler] [GetAll] error getting all products, err: ", err.Error())
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, products)
}

func (h *VariantHandler) GetByID(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		log.Println("[VariantHandler] [GetByID] error converting id to int, err: ", err.Error())
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	product, err := h.usecase.GetByID(id)
	if err != nil {
		log.Println("[VariantHandler] [GetByID] error getting product by id, err: ", err.Error())
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, product)
}

func (h *VariantHandler) Create(c echo.Context) error {
	product := new(domain.Variant)
	if err := c.Bind(&product); err != nil {
		log.Println("[VariantHandler] [Create] error binding product, err: ", err.Error())
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	product, err := h.usecase.Create(product)
	if err != nil {
		log.Println("[VariantHandler] [Create] error creating product, err: ", err.Error())
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, product)
}

func (h *VariantHandler) Update(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		log.Println("[VariantHandler] [Update] error converting id to int, err: ", err.Error())
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	product := new(domain.Variant)
	if err := c.Bind(&product); err != nil {
		log.Println("[VariantHandler] [Update] error binding product, err: ", err.Error())
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	product, err = h.usecase.Update(id, product)
	if err != nil {
		log.Println("[VariantHandler] [Update] error updating product, err: ", err.Error())
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, product)
}

func (h *VariantHandler) Delete(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		log.Println("[VariantHandler] [Delete] error converting id to int, err: ", err.Error())
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	err = h.usecase.Delete(id)
	if err != nil {
		log.Println("[VariantHandler] [Delete] error deleting product, err: ", err.Error())
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, "Variant deleted")
}
