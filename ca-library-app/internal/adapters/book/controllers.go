package book

import "ca-library-app/internal/book"

type handler struct {
	bookService book.Service
}

func New