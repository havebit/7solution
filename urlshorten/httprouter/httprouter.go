package httprouter

import (
	"encoding/json"
	"log/slog"
	"net/http"
)

type Context interface {
	Bind(any) error
	Param(string) string
	Status(int)
	JSON(int, any)
	Redirect(int, string)
}

func NewHandler(handler func(Context)) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		c := NewContext(w, r)
		handler(c)
	})
}

type context struct {
	w http.ResponseWriter
	r *http.Request
}

func NewContext(w http.ResponseWriter, r *http.Request) *context {
	return &context{w, r}
}

func (c *context) Bind(v any) error {
	return json.NewDecoder(c.r.Body).Decode(v)
}

func (c *context) Param(k string) string {
	return c.r.PathValue(k)
}

func (c *context) Status(code int) {
	c.w.WriteHeader(code)
}

func (c *context) JSON(code int, v any) {
	if err := json.NewEncoder(c.w).Encode(v); err != nil {
		slog.Error(err.Error())
	}
}

func (c *context) Redirect(code int, url string) {
	http.Redirect(c.w, c.r, url, code)
}
