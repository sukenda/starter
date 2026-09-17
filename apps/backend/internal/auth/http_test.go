package auth

import (
	"net/http/httptest"
	"testing"
	"time"

	"github.com/gofiber/fiber/v3"
)

func TestWebCookieScopes(t *testing.T) {
	app := fiber.New()
	h := NewHandler(nil, true)
	app.Get("/cookies", func(c fiber.Ctx) error {
		return h.setWebCookies(c, Tokens{
			RefreshToken:     "refresh-secret",
			RefreshExpiresAt: time.Now().Add(time.Hour),
		})
	})

	response, err := app.Test(httptest.NewRequest("GET", "/cookies", nil))
	if err != nil {
		t.Fatalf("request failed: %v", err)
	}

	cookies := response.Cookies()
	if len(cookies) != 2 {
		t.Fatalf("expected 2 cookies, got %d", len(cookies))
	}

	byName := make(map[string]struct {
		path     string
		httpOnly bool
		secure   bool
	})
	for _, cookie := range cookies {
		byName[cookie.Name] = struct {
			path     string
			httpOnly bool
			secure   bool
		}{cookie.Path, cookie.HttpOnly, cookie.Secure}
	}

	refresh := byName[refreshCookie]
	if refresh.path != refreshCookiePath || !refresh.httpOnly || !refresh.secure {
		t.Fatalf("unexpected refresh cookie attributes: %+v", refresh)
	}
	csrf := byName[csrfCookie]
	if csrf.path != csrfCookiePath || csrf.httpOnly || !csrf.secure {
		t.Fatalf("unexpected csrf cookie attributes: %+v", csrf)
	}
}

func TestClearWebCookiesUsesOriginalScopes(t *testing.T) {
	app := fiber.New()
	h := NewHandler(nil, true)
	app.Get("/cookies", func(c fiber.Ctx) error {
		h.clearWebCookies(c)
		return c.SendStatus(fiber.StatusNoContent)
	})

	response, err := app.Test(httptest.NewRequest("GET", "/cookies", nil))
	if err != nil {
		t.Fatalf("request failed: %v", err)
	}

	cookies := response.Cookies()
	if len(cookies) != 2 {
		t.Fatalf("expected 2 clearing cookies, got %d", len(cookies))
	}
	paths := map[string]string{}
	for _, cookie := range cookies {
		paths[cookie.Name] = cookie.Path
		if cookie.MaxAge >= 0 {
			t.Fatalf("expected %s to be expired, max age %d", cookie.Name, cookie.MaxAge)
		}
	}
	if paths[refreshCookie] != refreshCookiePath {
		t.Fatalf("refresh clear path = %q", paths[refreshCookie])
	}
	if paths[csrfCookie] != csrfCookiePath {
		t.Fatalf("csrf clear path = %q", paths[csrfCookie])
	}
}
