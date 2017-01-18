package loket

import (
	"github.com/mataharimall/micro/service"
	"github.com/mataharimall/micro/apps/loket/handlers"
	"github.com/mataharimall/micro/middleware"
	"github.com/labstack/echo"
)

type LoketRoute struct{}

func init() {
	service.ServiceManager.Register("route.loket", &LoketRoute{})
}

func (l *LoketRoute) SetRoute(s *echo.Echo) *echo.Echo {
	s.Use(middleware.Logger())
	s.Get("/loket/event", handlers.GetEventList)
	return s
}
