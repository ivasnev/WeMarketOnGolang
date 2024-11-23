package dto

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginResponse struct {
	Token string `json:"token"`
}

type UserResponse struct {
	ID    int32  `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
	Role  string `json:"role"`
}

type UserSessionResponse struct {
	ID        int32  `json:"id"`
	UserID    int32  `json:"user_id"`
	IPAddress string `json:"ip_address"`
	StartDate string `json:"start_date"`
	EndDate   string `json:"end_date"`
}
