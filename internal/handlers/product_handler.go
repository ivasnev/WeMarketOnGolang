package handlers

import (
	"WeMarketOnGolang/internal/dto"
	"WeMarketOnGolang/internal/models"
	"WeMarketOnGolang/internal/services/products" // Путь к пакету с интерфейсом ProductService
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// ProductHandler представляет обработчики для продуктов.
type ProductHandler struct {
	ProductService products.ProductServiceInterface
}

// NewProductHandler создает новый экземпляр ProductHandler.
func NewProductHandler(service products.ProductServiceInterface) *ProductHandler {
	return &ProductHandler{ProductService: service}
}

// CreateProduct создает новый продукт.
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
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create product"})
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

// GetProductByID возвращает продукт по ID.
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

// GetAllProducts возвращает все продукты.
func (h *ProductHandler) GetAllProducts(c *gin.Context) {
	allProducts, err := h.ProductService.GetAllProducts()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch allProducts"})
		return
	}

	response := []dto.ProductResponseDTO{}
	for _, product := range allProducts {
		response = append(response, dto.ProductResponseDTO{
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

// UpdateProduct обновляет продукт по ID.
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
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update product"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Product updated successfully"})
}

// DeleteProduct удаляет продукт по ID.
func (h *ProductHandler) DeleteProduct(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	if err := h.ProductService.DeleteProduct(int32(id)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete product"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Product deleted successfully"})
}
