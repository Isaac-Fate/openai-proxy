package handlers

import "github.com/gin-gonic/gin"

func WelcomeHandler(ctx *gin.Context) {
	ctx.HTML(200, "index.html", gin.H{
		"Title":   "OpenAI Proxy",
		"Content": "We will hanlde the OpenAI API requests for you :)",
	})
}
