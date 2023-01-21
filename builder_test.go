package broid

import (
	"github.com/stretchr/testify/require"
	"net/http"
	"testing"
)

func TestFletcher8(t *testing.T) {
	req := require.New(t)

	req.Zero(Fletcher8(""))

	req.NotEqual(Fletcher8("aa"), Fletcher8("ab"))
	req.NotEqual(Fletcher8("aa"), Fletcher8("bb"))
	req.NotEqual(Fletcher8("ab"), Fletcher8("ba"))
}

func TestHeaderFieldFunc(t *testing.T) {
	f := HeaderFieldFunc("Accept")

	var req http.Request
	req.Header = make(map[string][]string)
	require.Equal(t, uint8(255), f(&req))

	req.Header.Set("Accept", "")
	require.Equal(t, uint8(255), f(&req))

	req.Header.Set("Accept", "aaa")
	v1 := f(&req)
	req.Header.Set("Accept", "bbb")
	v2 := f(&req)
	require.NotEqual(t, v1, v2)
}

func TestCookieFieldFunc(t *testing.T) {
	f := CookieFieldFunc("test")

	var req http.Request
	require.Equal(t, uint8(255), f(&req))

	req.Header = make(map[string][]string)
	req.AddCookie(&http.Cookie{
		Name:  "test",
		Value: "aaa",
	})
	v1 := f(&req)
	req.Header = make(map[string][]string)
	req.AddCookie(&http.Cookie{
		Name:  "test",
		Value: "bbb",
	})
	v2 := f(&req)
	require.NotEqual(t, v1, v2)
}

func TestCookieNumberFieldFunc(t *testing.T) {
	f := CookieNumberFieldFunc("test")

	var req http.Request
	require.Equal(t, uint8(255), f(&req))

	req.Header = make(map[string][]string)
	req.AddCookie(&http.Cookie{
		Name:  "test",
		Value: "aaa",
	})
	require.Equal(t, uint8(255), f(&req))

	req.Header = make(map[string][]string)
	req.AddCookie(&http.Cookie{
		Name:  "test",
		Value: "123456",
	})
	require.Equal(t, uint8(36), f(&req))
}

func TestBuild(t *testing.T) {
	b := NewDefaultBrowserIDBuilder()

	var req http.Request
	req.Header = make(map[string][]string)
	require.Equal(t, "ffffffff", b.Build(&req).String())

	req.Header.Set("Accept", "aaa")
	req.Header.Set("User-Agent", "uuu")
	v1 := b.Build(&req).String()
	req.Header.Set("Accept", "bbb")
	req.Header.Set("User-Agent", "user-agent")
	v2 := b.Build(&req).String()
	require.NotEqual(t, v1, v2)
}
