package middleware

import (
	"first/go_web/response"
	"fmt"
	"github.com/gin-gonic/gin"
)

func RecoveryMiddleware() gin.HandlerFunc{
	return func(ctx *gin.Context){
		defer func(){
			if err := recover();err != nil {
				response.Fail(ctx,nil,fmt.Sprint(err))
			}
		}()
		ctx.Next()
	}
}