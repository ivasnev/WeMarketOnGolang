package handlers

import (
	"WeMarketOnGolang/internal/dto"
	"WeMarketOnGolang/internal/models"
	"WeMarketOnGolang/internal/services/products"
	"WeMarketOnGolang/internal/utils"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"net/http"
	"strconv"
)

type ProductHandler struct {
	ProductService products.ProductServiceInterface
}

func NewProductHandler(service products.ProductServiceInterface) *ProductHandler {
	return &ProductHandler{ProductService: service}
}

// CreateProduct создает новый продукт
// @Summary Создать продукт
// @Description Создает новый продукт на основе предоставленных данных
// @Tags v1/products
// @Accept json
// @Produce json
// @Param product body dto.CreateProductDTO true "Данные продукта"
// @Success 201 {object} dto.ProductResponseDTO
// @Failure 400 {object} map[string]interface{} "Неверные данные"
// @Failure 500 {object} map[string]interface{} "Ошибка сервера"
// @Security BearerAuth
// @Router /v1/products [post]
func (h *ProductHandler) CreateProduct(c *gin.Context) {
	var dtoObj dto.CreateProductDTO
	if err := c.ShouldBindJSON(&dtoObj); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input", "details": err.Error()})
		return
	}

	product := &models.Product{
		Name:               dtoObj.Name,
		Price:              dtoObj.Price,
		CategoryID:         dtoObj.CategoryID,
		Stock:              dtoObj.Stock,
		AvailabilityStatus: dtoObj.AvailabilityStatus,
		ManufacturerID:     dtoObj.ManufacturerID,
	}

	if dtoObj.Description != nil {
		product.Description = dtoObj.Description
	}
	if dtoObj.ImageURL != nil {
		product.ImageURL = dtoObj.ImageURL
	}
	if dtoObj.Sku != nil {
		product.Sku = dtoObj.Sku
	}
	if dtoObj.Weight != nil {
		product.Weight = dtoObj.Weight
	}

	if err := h.ProductService.CreateProduct(product); err != nil {
		statusCode, response := utils.HandleDBError(err)
		c.JSON(statusCode, response)
		return
	}

	c.JSON(http.StatusCreated, dto.ProductResponseDTO{
		ID:                 product.ID,
		Name:               product.Name,
		Description:        product.Description,
		Price:              product.Price,
		CategoryID:         product.CategoryID,
		Stock:              product.Stock,
		ImageURL:           product.ImageURL,
		AddedDate:          product.AddedDate,
		Sku:                product.Sku,
		Weight:             product.Weight,
		AvailabilityStatus: product.AvailabilityStatus,
		ManufacturerID:     product.ManufacturerID,
	})
}

// GetProductByID возвращает продукт по ID
// @Summary Получить продукт по ID
// @Description Возвращает продукт на основе переданного идентификатора
// @Tags v1/products
// @Accept json
// @Produce json
// @Param id path int true "ID продукта"
// @Success 200 {object} dto.ProductResponseDTO
// @Failure 400 {object} map[string]interface{} "Неверный ID"
// @Failure 404 {object} map[string]interface{} "Продукт не найден"
// @Security BearerAuth
// @Router /v1/products/{id} [get]
func (h *ProductHandler) GetProductByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	product, err := h.ProductService.GetProductByID(int32(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Product not found"})
		return
	}

	c.JSON(http.StatusOK, dto.ProductResponseDTO{
		ID:                 product.ID,
		Name:               product.Name,
		Description:        product.Description,
		Price:              product.Price,
		CategoryID:         product.CategoryID,
		Stock:              product.Stock,
		ImageURL:           product.ImageURL,
		AddedDate:          product.AddedDate,
		Sku:                product.Sku,
		Weight:             product.Weight,
		AvailabilityStatus: product.AvailabilityStatus,
		ManufacturerID:     product.ManufacturerID,
	})
}

// GetAllProducts возвращает все продукты
// @Summary Получить все продукты
// @Description Возвращает список всех продуктов с поддержкой фильтрации и пагинации
// @Tags v1/products
// @Accept json
// @Produce json
// @Param page query int false "Номер страницы"
// @Param page_size query int false "Размер страницы"
// @Param name query string false "Название продукта"
// @Param min_price query number false "Минимальная цена"
// @Param max_price query number false "Максимальная цена"
// @Success 200 {object} dto.ProductResponseDTOWithPagination
// @Failure 422 {object} map[string]interface{} "Ошибка валидации"
// @Failure 500 {object} map[string]interface{} "Ошибка сервера"
// @Security BearerAuth
// @Router /v1/products [get]
func (h *ProductHandler) GetAllProducts(c *gin.Context) {

	filter := &dto.ProductFilter{}
	if err := c.ShouldBindQuery(&filter); err != nil {
		var errorMessages []string

		if errs, ok := err.(validator.ValidationErrors); ok {
			for _, e := range errs {
				errorMessages = append(errorMessages, fmt.Sprintf("Field '%s' failed validation, condition: %s", e.Field(), e.ActualTag()))
			}
		} else {
			errorMessages = append(errorMessages, err.Error())
		}

		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"error":   "Validation failed",
			"details": errorMessages,
			"code":    http.StatusUnprocessableEntity,
		})
		return
	}

	allProducts, total, err := h.ProductService.GetAllProducts(filter)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch products"})
		return
	}

	response := dto.ProductResponseDTOWithPagination{
		Products: []dto.ProductResponseDTO{},
		Total:    total,
		PageSize: filter.PageSize,
		Page:     filter.Page,
	}

	for _, product := range allProducts {
		response.Products = append(response.Products, dto.ProductResponseDTO{
			ID:                 product.ID,
			Name:               product.Name,
			Description:        product.Description,
			Price:              product.Price,
			CategoryID:         product.CategoryID,
			Stock:              product.Stock,
			ImageURL:           product.ImageURL,
			AddedDate:          product.AddedDate,
			Sku:                product.Sku,
			Weight:             product.Weight,
			AvailabilityStatus: product.AvailabilityStatus,
			ManufacturerID:     product.ManufacturerID,
		})
	}

	c.JSON(http.StatusOK, response)
}

// UpdateProduct обновляет продукт
// @Summary Обновить продукт
// @Description Обновляет данные продукта на основе переданных данных
// @Tags v1/products
// @Accept json
// @Produce json
// @Param id path int true "ID продукта"
// @Param product body dto.UpdateProductDTO true "Данные для обновления продукта"
// @Success 200 {object} map[string]interface{} "Продукт успешно обновлен"
// @Failure 400 {object} map[string]interface{} "Неверный ID или данные"
// @Failure 500 {object} map[string]interface{} "Ошибка сервера"
// @Security BearerAuth
// @Router /v1/products/{id} [put]
func (h *ProductHandler) UpdateProduct(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	var dtoObj dto.UpdateProductDTO
	if err := c.ShouldBindJSON(&dtoObj); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input", "details": err.Error()})
		return
	}

	var product models.Product

	if dtoObj.Name != nil {
		product.Name = *dtoObj.Name
	}
	if dtoObj.Description != nil {
		product.Description = dtoObj.Description
	}
	if dtoObj.Price != nil {
		product.Price = *dtoObj.Price
	}
	if dtoObj.CategoryID != nil {
		product.CategoryID = *dtoObj.CategoryID
	}
	if dtoObj.Stock != nil {
		product.Stock = *dtoObj.Stock
	}
	if dtoObj.ImageURL != nil {
		product.ImageURL = dtoObj.ImageURL
	}
	if dtoObj.Sku != nil {
		product.Sku = dtoObj.Sku
	}
	if dtoObj.Weight != nil {
		product.Weight = dtoObj.Weight
	}
	if dtoObj.AvailabilityStatus != nil {
		product.AvailabilityStatus = *dtoObj.AvailabilityStatus
	}
	if dtoObj.ManufacturerID != nil {
		product.ManufacturerID = dtoObj.ManufacturerID
	}

	if err := h.ProductService.UpdateProduct(int32(id), &product); err != nil {
		statusCode, response := utils.HandleDBError(err)
		c.JSON(statusCode, response)
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Product updated successfully"})
}

// DeleteProduct удаляет продукт
// @Summary Удалить продукт
// @Description Удаляет продукт по его идентификатору
// @Tags v1/products
// @Accept json
// @Produce json
// @Param id path int true "ID продукта"
// @Success 200 {object} map[string]interface{} "Продукт успешно удален"
// @Failure 400 {object} map[string]interface{} "Неверный ID"
// @Failure 500 {object} map[string]interface{} "Ошибка сервера"
// @Security BearerAuth
// @Router /v1/products/{id} [delete]
func (h *ProductHandler) DeleteProduct(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	if err := h.ProductService.DeleteProduct(int32(id)); err != nil {
		statusCode, response := utils.HandleDBError(err)
		c.JSON(statusCode, response)
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Product deleted successfully"})
}
