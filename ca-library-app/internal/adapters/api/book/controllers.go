package book

import (
	"ca-library-app/internal/adapters"
	"ca-library-app/internal/book"

	//"context"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

const (
	bookURL  = "/books/:book_id"
	booksURL = "/books"
)

type handler struct {
	bookService book.Service
}

func (h *handler) Register(router *httprouter.Router) {
	router.GET(booksURL, h.GetAllBooks)
}

func NewHandler(service book.Service) adapters.Handler {
	return &handler{bookService: service}
}

func (h *handler) GetAllBooks(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	//books := h.bookService.GetAllBooks(context.Background(), 0, 0)
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("books"))
}
