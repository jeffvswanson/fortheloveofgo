package book

import "testing"

func TestBook(t *testing.T) {
	_ = Book{
		Title:  "Spark Joy",
		Author: []string{"Marie Kondo"},
		Price:  1199,
	}
}
