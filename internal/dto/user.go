package dto

// DTO для создания нового пользователя
type CreateUserDTO struct {
	Email    string  `json:"email" binding:"required,email"`
	Password string  `json:"password" binding:"required,min=6"`
	Name     string  `json:"name" binding:"required"`
	Phone    *string `json:"phone" binding:"required"`
	Address  *string `json:"address"`
}

// DTO для обновления данных пользователя
type UpdateUserDTO struct {
	Name    *string `json:"name"`
	Phone   *string `json:"phone"`
	Address *string `json:"address"`
}

// DTO для ответа с данными пользователя
type UserResponseDTO struct {
	ID      int32   `json:"id"`
	Email   string  `json:"email"`
	Name    string  `json:"name"`
	Phone   *string `json:"phone"`
	Address *string `json:"address"`
	RoleID  int32   `json:"role_id"`
}
