package core

import (
	"os"
	"strconv"
	"time"

	"myapp/internal/views/fragments"
	"myapp/internal/views/pages"

	"github.com/a-h/templ"
	"github.com/labstack/echo/v5"
	"github.com/pocketbase/pocketbase/apis"
)

type PageHandler func() templ.Component

const DEFAULT_CACHE_DURATION = 24 * time.Hour

// This is where you define the routes of your application
// You can also define middlewares here
func MountRoutes(app *App, router *echo.Echo) error {
	// Pages
	router.GET("/", renderTempl(pages.IndexPage))

	// Fragments
	router.GET("/frag/time", renderTempl(func() templ.Component {
		return fragments.Time(time.Now())
	}))

	// Static content
	router.GET("/*", apis.StaticDirectoryHandler(os.DirFS("./public"), false), cached(DEFAULT_CACHE_DURATION))

	return nil
}

func renderTempl(h PageHandler) echo.HandlerFunc {
	return func(c echo.Context) error {
		return h().Render(c.Request().Context(), c.Response().Writer)
	}
}

// cached is a middleware that sets the Cache-Control header for successful responses.
// It takes a maxAge duration which is used to set the max-age directive in the header.
func cached(maxAge time.Duration) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			if err := next(c); err != nil {
				c.Error(err)
			} else {
				if c.Response().Status >= 200 && c.Response().Status < 300 {
					seconds := strconv.Itoa(int(maxAge.Seconds()))
					c.Response().Header().Set("Cache-Control", "public, max-age="+seconds)
				}
			}
			return nil
		}
	}
}
