package middleware

import (
	"github.com/XCPCBoard/utils/errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

func ErrorHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next() // 先调用c.Next()执行后面的中间件
		// 所有中间件及router处理完毕后从这里开始执行
		// 检查c.Errors中是否有错误
		for _, e := range c.Errors {
			err := e.Err
			// 若是自定义的错误则将code、msg返回
			if myErr, ok := err.(*errors.MyError); ok {
				c.AbortWithError(http.StatusNotAcceptable, myErr)
			} else {
				// 若非自定义错误则返回详细错误信息err.Error()
				c.AbortWithError(http.StatusNotAcceptable, errors.GetError(errors.INNER_ERROR, err.Error()))
			}
			return // 检查一个错误就行
		}
	}
}
