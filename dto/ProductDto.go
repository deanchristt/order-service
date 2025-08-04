package dto

type ProductUpdateDto struct {
	ID          int    `json:"-"`
	Title       string `json:"title" form:"title" binding:"required"`
	Description string `json:"description" form:"description" binding:"required"`
	UserId      int    `json:"userId" form:"userId,omitempty"`
}

type ProductCreateDto struct {
	Title       string `json:"title" form:"title" binding:"required"`
	Description string `json:"description" form:"description" binding:"required"`
	UserId      int    `json:"userId" form:"userId,omitempty"`
}
