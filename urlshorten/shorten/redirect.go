package shorten

import (
	"net/http"

	"github.com/pallat/urlshorten/httprouter"
)

type originalURLGetter interface {
	OriginalURL(shortenURL string) (string, error)
}

type redirectHandler struct {
	storage originalURLGetter
}

func NewRedirectHandler(storage originalURLGetter) *redirectHandler {
	return &redirectHandler{storage: storage}
}

func (handler *redirectHandler) Handler(c httprouter.Context) {
	shortenURL := c.Param("shorturl")
	var originalURL string

	originalURL, err := handler.storage.OriginalURL(shortenURL)
	if err != nil {
		c.Status(http.StatusNotFound)
		return
	}

	c.Redirect(http.StatusFound, originalURL)
}
