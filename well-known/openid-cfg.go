package wellknown

import (
	"github.com/labstack/echo"
	"net/http"
)

// Index lists all wellknown endpoints
func Index(c echo.Context) error {
	return c.String(http.StatusOK, "Index wellknown")
}
