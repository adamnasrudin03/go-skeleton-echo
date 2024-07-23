package router

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

type Routes struct {
	HttpServer *echo.Echo
}

func NewRoutes() *Routes {
	r := &Routes{
		HttpServer: echo.New(),
	}

	r.HttpServer.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: `{"time":"${time_rfc3339_nano}","id":"${id}","remote_ip":"${remote_ip}","host":"${host}",` +
			`"method":"${method}","uri":"${uri}","status":${status},"error":"${error}",` +
			`"latency_human":"${latency_human}"}` + "\n",
	}))
	r.HttpServer.Use(middleware.RequestID())
	r.HttpServer.Use(middleware.Recover())

	return r
}

func (r *Routes) Run(addr string) error {
	return r.HttpServer.Start(addr)
}
