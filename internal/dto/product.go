package dto

import (
	"time"
)

// CreateProductDTO структура для создания продукта.
type CreateProductDTO struct {
	Name               string   `json:"name" binding:"required"`
	Description        *string  `json:"description"`
	Price              float64  `json:"price" binding:"required"`
	CategoryID         int32    `json:"category_id" binding:"required"`
	Stock              int32    `json:"stock" binding:"required"`
	ImageURL           *string  `json:"image_url"`
	Sku                *string  `json:"sku"`
	Weight             *float64 `json:"weight"`
	AvailabilityStatus int32    `json:"availability_status" binding:"required"`
	ManufacturerID     *int32   `json:"manufacturer_id"`
}

// UpdateProductDTO структура для обновления продукта.
type UpdateProductDTO struct {
	Name               *string  `json:"name"`
	Description        *string  `json:"description"`
	Price              *float64 `json:"price"`
	CategoryID         *int32   `json:"category_id"`
	Stock              *int32   `json:"stock"`
	ImageURL           *string  `json:"image_url"`
	Sku                *string  `json:"sku"`
	Weight             *float64 `json:"weight"`
	AvailabilityStatus *int32   `json:"availability_status"`
	ManufacturerID     *int32   `json:"manufacturer_id"`
}

// ProductResponseDTO структура для ответа с информацией о продукте.
type ProductResponseDTO struct {
	ID                 int32      `json:"id"`
	Name               string     `json:"name"`
	Description        *string    `json:"description"`
	Price              float64    `json:"price"`
	CategoryID         int32      `json:"category_id"`
	Stock              int32      `json:"stock"`
	ImageURL           *string    `json:"image_url"`
	AddedDate          *time.Time `json:"added_date"`
	Sku                *string    `json:"sku"`
	Weight             *float64   `json:"weight"`
	AvailabilityStatus int32      `json:"availability_status"`
	ManufacturerID     *int32     `json:"manufacturer_id"`
}
