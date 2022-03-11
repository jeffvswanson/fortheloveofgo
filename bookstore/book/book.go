package book

// Book represents information about a book.
type Book struct {
	Author          []string
	Copies          int
	DiscountPercent uint // 10 represents a 10% discount
	Edition         int
	Featured        bool // Book is selected for Read of the Month
	Price           int  // Price in cents
	SeriesOrder     int
	Title           string
}

// NetPrice calculates the prices of a book given its DiscountPercent
func (b Book) NetPrice() int {
	return b.Price - ((b.Price * int(b.DiscountPercent)) / 100)
}
