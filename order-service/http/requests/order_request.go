package requests

type CreateOrderRequest struct {
	Id     uint    `json:"id"`
	Total  float64 `json:"total" binding:"required"`
	UserID uint    `json:"user_id" binding:"required"`
}
