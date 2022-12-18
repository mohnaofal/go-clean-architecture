package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/mohnaofal/go-clean-architecture/app/models"
	"github.com/mohnaofal/go-clean-architecture/app/services"
)

type ProductsHandler struct {
	product services.Products
}

func NewProductsHandler(product services.Products) ProductsHandler {
	return ProductsHandler{product: product}
}

func (h *ProductsHandler) Mount(c *echo.Group) {
	c.POST("", h.Create)
	c.PUT("", h.Update)
	c.GET("", h.View)
	c.GET("/:code", h.Detail)
	c.DELETE("/:code", h.Detele)
}

func (h *ProductsHandler) Create(c echo.Context) error {
	ctx := c.Request().Context()
	form := new(models.Product)
	if err := c.Bind(form); err != nil {
		return c.JSON(http.StatusUnprocessableEntity, map[string]interface{}{
			"success": false,
			"message": err.Error(),
		})
	}

	data, err := h.product.Create(ctx, form)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"success": false,
			"message": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, data)
}

func (h *ProductsHandler) Update(c echo.Context) error {
	ctx := c.Request().Context()
	form := new(models.Product)

	if err := c.Bind(form); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"success": false,
			"message": err.Error(),
		})
	}

	data, err := h.product.Update(ctx, form)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"success": false,
			"message": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, data)
}

func (h *ProductsHandler) View(c echo.Context) error {
	ctx := c.Request().Context()

	data, err := h.product.View(ctx)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"success": false,
			"message": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, data)
}

func (h *ProductsHandler) Detail(c echo.Context) error {
	ctx := c.Request().Context()

	productCode := c.Param("code")
	data, err := h.product.Detail(ctx, productCode)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"success": false,
			"message": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, data)
}

func (h *ProductsHandler) Detele(c echo.Context) error {
	ctx := c.Request().Context()

	productCode := c.Param("code")
	if err := h.product.Delete(ctx, productCode); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"success": false,
			"message": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"success": true,
	})
}
