package helpers

import (
	"github.com/gin-gonic/gin"
)

func DefaultApiResponseObject(c *gin.Context, httpStatus int, data gin.H) {
	c.JSON(httpStatus, data)
}
