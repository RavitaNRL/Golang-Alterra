package payload

type CreateItemRequest struct {
	Name        string `json:"name" validate:"required"`
	Description string `json:"description" validate:"required"`
	Stock       uint   `json:"stock" validate:"required"`
	Price       uint   `json:"price" validate:"required"`
	CategoryID  uint   `json:"category_id" validate:"required"`
}

type CategoryRequest struct {
	Name string `json:"name" validate:"required"`
}

type CreateUserRequest struct {
	Name     string `json:"name" validate:"required"`
	Email    string `json:"email" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type LoginRequest struct {
	Email    string `json:"email" form:"email" validate:"required,email"`
	Password string `json:"password" form:"password" validate:"required"`
}
