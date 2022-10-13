package products

import (
	"log"
	"net/http"
	"strconv"

	"github.com/bpradana/devcamp_pe_bintangpradanaerlanggaputra/pkg/domain"
	"github.com/labstack/echo/v4"
)

type ProductHandler struct {
	usecase domain.ProductUsecase
}

func NewHandler(e *echo.Group, usecase domain.ProductUsecase) {
	h := &ProductHandler{
		usecase: usecase,
	}

	// Routes
	e.GET("/products", h.GetAll)
	e.GET("/products/:id", h.GetByID)
	e.POST("/products", h.Create)
	e.PUT("/products/:id", h.Update)
	e.DELETE("/products/:id", h.Delete)
}

func (h *ProductHandler) GetAll(c echo.Context) error {
	products, err := h.usecase.GetAll()
	if err != nil {
		log.Println("[ProductHandler] [GetAll] error getting all products, err: ", err.Error())
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, products)
}

func (h *ProductHandler) GetByID(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		log.Println("[ProductHandler] [GetByID] error converting id to int, err: ", err.Error())
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	product, err := h.usecase.GetByID(id)
	if err != nil {
		log.Println("[ProductHandler] [GetByID] error getting product by id, err: ", err.Error())
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, product)
}

func (h *ProductHandler) Create(c echo.Context) error {
	product := new(domain.Product)
	if err := c.Bind(&product); err != nil {
		log.Println("[ProductHandler] [Create] error binding product, err: ", err.Error())
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	product, err := h.usecase.Create(product)
	if err != nil {
		log.Println("[ProductHandler] [Create] error creating product, err: ", err.Error())
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, product)
}

func (h *ProductHandler) Update(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		log.Println("[ProductHandler] [Update] error converting id to int, err: ", err.Error())
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	product := new(domain.Product)
	if err := c.Bind(&product); err != nil {
		log.Println("[ProductHandler] [Update] error binding product, err: ", err.Error())
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	product, err = h.usecase.Update(id, product)
	if err != nil {
		log.Println("[ProductHandler] [Update] error updating product, err: ", err.Error())
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, product)
}

func (h *ProductHandler) Delete(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		log.Println("[ProductHandler] [Delete] error converting id to int, err: ", err.Error())
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	err = h.usecase.Delete(id)
	if err != nil {
		log.Println("[ProductHandler] [Delete] error deleting product, err: ", err.Error())
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, "Product deleted")
}
