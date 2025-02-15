package dtos


type WishlistRequestDto struct {
	Title       string `json:"title" binding:"required_if=IsUpdate false"`
	Description string `json:"description" binding:"required_if=IsUpdate false"`
	EventDate   string `json:"event_date" binding:"required_if=IsUpdate false"`
	IsUpdate    bool    `json:"-"`
}
