package broid

import (
	"net/http"
	"strconv"
)

// FieldFunc is a function that computes a single BrowserID field.
type FieldFunc func(req *http.Request) uint8

// BrowserIDBuilder contains functions for computing a BrowserID.
type BrowserIDBuilder struct {
	fields []FieldFunc
}

// NewDefaultBrowserIDBuilder returns the default BrowserIDBuilder.
// It contains functions for computing a BrowserID based on the
// "User-Agent", "Accept", "Accept-Encoding" and "Accept-Language" headers.
func NewDefaultBrowserIDBuilder() *BrowserIDBuilder {
	var b BrowserIDBuilder

	b.AddField(HeaderFieldFunc("User-Agent"))
	b.AddField(HeaderFieldFunc("Accept"))
	b.AddField(HeaderFieldFunc("Accept-Encoding"))
	b.AddField(HeaderFieldFunc("Accept-Language"))

	return &b
}

// AddField appends a FieldFunc to a BrowserIDBuilder.
func (b *BrowserIDBuilder) AddField(f FieldFunc) {
	b.fields = append(b.fields, f)
}

// Build returns a BrowserID computed for a given Request.
func (b *BrowserIDBuilder) Build(req *http.Request) BrowserID {
	id := make([]uint8, 0, len(b.fields))

	for _, f := range b.fields {
		val := f(req)
		id = append(id, val)
	}

	return id
}

// HeaderFieldFunc computes a BrowserID field from a header value.
func HeaderFieldFunc(key string) FieldFunc {
	return func(req *http.Request) uint8 {
		val := req.Header.Get(key)
		if val == "" {
			return 255
		}

		return Fletcher8(val)
	}
}

// CookieFieldFunc computes a BrowserID field from a string stored in cookies.
func CookieFieldFunc(key string) FieldFunc {
	return func(req *http.Request) uint8 {
		c, err := req.Cookie(key)
		if err != nil {
			return 255
		}

		return Fletcher8(c.Value)
	}
}

// CookieNumberFieldFunc computes a BrowserID field from a number stored in cookies.
func CookieNumberFieldFunc(key string) FieldFunc {
	return func(req *http.Request) uint8 {
		c, err := req.Cookie(key)
		if err != nil {
			return 255
		}

		val, err := strconv.ParseInt(c.Value, 10, 64)
		if err != nil {
			return 255
		}

		return uint8(val % 255)
	}
}

// Fletcher8 calculates hash sum of the given string.
// It can be used to compute a BrowserID field based on a string value.
func Fletcher8(val string) uint8 {
	var s1, s2 uint32

	for _, b := range []byte(val) {
		s1 += uint32(b)
		s2 += s1
	}

	s1 %= 15
	s2 %= 15

	return uint8((s2 << 4) | s1)
}
