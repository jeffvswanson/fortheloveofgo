package book

// Book represents information about a book.
type Book struct {
	Title           string
	Author          string
	Copies          int
	Edition         int
	Price           int // Price in cents
	SeriesOrder     int
	DiscountPercent uint // 10 represents a 10% discount
}
