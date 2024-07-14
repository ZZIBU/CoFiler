package utils

import (
	"CoFiler/services/common/types"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GeneralResponse(c *gin.Context, code int, description string, errCode int, result interface{}) {
	c.JSON(code, &types.GeneralResponse{
		ResultCode:  code,
		Description: description,
		ErrCode:     errCode,
		Result:      result,
	})
}

func ResponseOK(c *gin.Context, response interface{}) {
	c.JSON(http.StatusOK, response)
}

func ResponseErr(c *gin.Context, err ...interface{}) {
	c.JSON(http.StatusInternalServerError, err)
}
