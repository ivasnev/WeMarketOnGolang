package dto

// InventoryStatusDTO представляет данные, которые будут переданы в запросах.
type InventoryStatusDTO struct {
	Name string `json:"name" binding:"required"`
}
