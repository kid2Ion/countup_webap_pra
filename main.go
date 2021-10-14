package main

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo"
)

func main() {
	e := echo.New()

	// e.Use(middleware.Logger())
	// e.Use(middleware.Recover())

	cnt := &counter{count: 0}

	e.GET("/", hello)
	e.GET("/count", cnt.countup)

	e.Logger.Fatal(e.Start(":1323"))
}

type counter struct {
	count int
}

func hello(c echo.Context) error {
	return c.String(http.StatusOK, "hello world")
}

func (c *counter) countup(ctx echo.Context) error {
	c.count += 1
	return ctx.String(http.StatusOK, strconv.Itoa(c.count))
}
