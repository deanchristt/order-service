package dto

type ProductUpdateDto struct {
	ID          int    `json:"-"`
	Title       string `json:"title" form:"title" binding:"required"`
	Description string `json:"description" form:"description" binding:"required"`
	CustomerId  int    `json:"customer_id" form:"customer_id,omitempty"`
}

type ProductCreateDto struct {
	Title       string `json:"title" form:"title" binding:"required"`
	Description string `json:"description" form:"description" binding:"required"`
	CustomerId  int    `json:"customer_id" form:"customer_id,omitempty"`
}
