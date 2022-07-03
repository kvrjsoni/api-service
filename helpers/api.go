package helpers

import (
	"github.com/gin-gonic/gin"
)

// default API response object used for all API calls
func DefaultApiResponseObject(c *gin.Context, httpStatus int, data gin.H) {
	c.JSON(httpStatus, data)
}
