package router

import (
	"Core/types"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (r *Router) GET(path string, handler ...gin.HandlerFunc) gin.IRoutes {
	return r.engine.GET(path, handler...)
}

func (r *Router) POST(path string, handler ...gin.HandlerFunc) gin.IRoutes {
	return r.engine.POST(path, handler...)
}

func (r *Router) PUT(path string, handler ...gin.HandlerFunc) gin.IRoutes {
	return r.engine.PUT(path, handler...)
}

func (r *Router) DELETE(path string, handler ...gin.HandlerFunc) gin.IRoutes {
	return r.engine.DELETE(path, handler...)
}

func (r *Router) GeneralResponse(c *gin.Context, code int, description string, errCode int, result interface{}) {
	c.JSON(code, &types.GeneralResponse{
		ResultCode:  code,
		Description: description,
		ErrCode:     errCode,
		Result:      result,
	})
}

func (r *Router) ResponseOK(c *gin.Context, response interface{}) {
	c.JSON(http.StatusOK, response)
}

func (r *Router) ResponseErr(c *gin.Context, err ...interface{}) {
	c.JSON(http.StatusInternalServerError, err)
}
