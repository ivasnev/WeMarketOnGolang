package dto

// CreateCategoryRequest represents the input for creating a category.
type CreateCategoryRequest struct {
	Name             string  `json:"name" binding:"required"`
	Description      *string `json:"description,omitempty"`
	ParentCategoryID *int32  `json:"parent_category_id,omitempty"`
}

// UpdateCategoryRequest represents the input for updating a category.
type UpdateCategoryRequest struct {
	Name             *string `json:"name,omitempty"`
	Description      *string `json:"description,omitempty"`
	ParentCategoryID *int32  `json:"parent_category_id,omitempty"`
}

// CategoryResponse represents the output returned to the client.
type CategoryResponse struct {
	ID               int32   `json:"id"`
	Name             string  `json:"name"`
	Description      *string `json:"description"`
	ParentCategoryID *int32  `json:"parent_category_id"`
}
