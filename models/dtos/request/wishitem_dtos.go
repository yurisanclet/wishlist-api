package dtos

type WishItemRequestDto struct {
	WishListID  string  `json:"wishlist_id" binding:"required_if=IsUpdate false"`
	Name        string  `json:"name" binding:"required_if=IsUpdate false"`
	Description string  `json:"description"`
	Link        string  `json:"link"`
	Price       float64 `json:"price"`
	Priority    string  `json:"priority" binding:"required_if=IsUpdate false"`
	IsBought    bool    `json:"is_bought"`
	IsUpdate    bool    `json:"-"`
}
