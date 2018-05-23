package main

import (
	"log"
	"net/http"

	"github.com/comail/colog"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	// "github.com/coreos/etcd/client"
)

const (
	listen_port    = ":1323"
	admin_username = "admin"
	admin_password = "changeme"
)

func indexHandler(c echo.Context) error {
	log.Printf("info: Get access")
	return c.String(http.StatusOK, "こんにちは")
}

func createRRHandler(c echo.Context) error {
	return c.String(http.StatusOK, "こんにちは")
}

func deleteRRHandler(c echo.Context) error {
	return c.String(http.StatusOK, "こんにちは")
}

func main() {
	colog.SetDefaultLevel(colog.LDebug)
	colog.SetMinLevel(colog.LTrace)
	colog.SetFormatter(&colog.StdFormatter{
		Colors: true,
		Flag:   log.Ldate | log.Ltime | log.Lshortfile,
	})
	colog.Register()

	e := echo.New()

	e.Use(middleware.BasicAuth(func(u, p string, c echo.Context) (bool, error) {
		if u == admin_username && p == admin_password {
			return true, nil
		}
		return false, nil
	}))

	e.GET("/", indexHandler)

	e.Start(listen_port)
}
