package shorten

import (
	"errors"
	"net/http"
	"testing"

	"github.com/pallat/urlshorten/httprouter"
)

type contextTest struct {
	httprouter.Context
	code int
}

func (c *contextTest) Bind(any) error {
	return nil
}

func (c *contextTest) JSON(code int, v any) {
	c.code = code
}

type storageStub struct{}

func (storageStub) Save(string, string) error {
	return errors.New("error test")
}

func TestHandler(t *testing.T) {
	handler := &shortenHandler{storage: &storageStub{}}
	ctx := &contextTest{}
	handler.Handler(ctx)

	if ctx.code != http.StatusInternalServerError {
		t.Errorf("it should be 500 but %v\n", ctx.code)
	}
}

type storageSpy struct {
	shortenURL string
}

func (s *storageSpy) Save(shortenURL, originalURL string) error {
	s.shortenURL = shortenURL
	return errors.New("end test")
}

func TestStorageSaveKeyValueCorrect(t *testing.T) {
	storage := &storageSpy{}
	ctx := &contextTest{}
	handler := &shortenHandler{storage: storage}

	handler.Handler(ctx)

	if len(storage.shortenURL) == 0 {
		t.Error("empty shorten")
	}
	if len(storage.shortenURL) > 6 {
		t.Error("too long")
	}
}
