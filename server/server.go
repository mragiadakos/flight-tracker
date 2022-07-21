package server

import (
	"github.com/labstack/echo/v4"
)

func NewServer(sorter Sorter) *Server {
	return &Server{sorter: sorter}
}

func (s *Server) Run(port string) {
	e := echo.New()
	e.HideBanner = true
	e.GET("/sort", s.sortHandler)
	e.Logger.Fatal(e.Start(port))
}
