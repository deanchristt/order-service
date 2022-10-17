package dto

type ProductUpdateDto struct {
	ID          int    `json:"id" form:"id" binding:"required"`
	Title       string `json:"title" form:"title" binding:"required"`
	Description string `json:"description" form:"description" binding:"required"`
	CustomerId  int    `json:"customer_id" form:"customer_id,omitempty" binding:"required"`
}

type ProductCreateDto struct {
	Title       string `json:"title" form:"title" binding:"required"`
	Description string `json:"description" form:"description" binding:"required"`
	CustomerId  int    `json:"customer_id" form:"customer_id,omitempty" binding:"required"`
}
