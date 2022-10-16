package dto

type ProductUpdateDto struct {
	ID          int    `json:"id" form:"id" binding:"required"`
	Title       string `json:"title" form:"title" binding:"required"`
	Description string `json:"description" form:"description" binding:"required"`
	UserId      int    `json:"user_id" form:"user_id,omitempty" binding:"required"`
}

type ProductCreateDto struct {
	Title       string `json:"title" form:"title" binding:"required"`
	Description string `json:"description" form:"description" binding:"required"`
	UserId      int    `json:"user_id" form:"user_id,omitempty" binding:"required"`
}
