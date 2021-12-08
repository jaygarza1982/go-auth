package controllers

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func Check() func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		session := sessions.Default(ctx)

		if username := session.Get("username"); username != nil {
			ctx.JSON(200, gin.H{"username": username})
			return
		}

		ctx.JSON(500, gin.H{"message": "Not authenticated"})

		return
	}
}
