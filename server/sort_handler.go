package server

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func (s *Server) sortHandler(c echo.Context) error {
	req := new(RequestSort)
	if err := c.Bind(req); err != nil {
		msg := "failed to unmarshal the body"
		return c.JSON(http.StatusBadRequest, &ResponseSort{
			Message: &msg,
			Status:  StatusFailure,
		})
	}

	path, err := s.sorter.Sort(req.Paths)
	if err != nil {
		msg := err.Error()
		return c.JSON(http.StatusBadRequest, &ResponseSort{
			Message: &msg,
			Status:  StatusFailure,
		})
	}

	return c.JSON(http.StatusOK, &ResponseSort{
		Path:   path,
		Status: StatusSuccess,
	})
}
