package Response

import "github.com/gin-gonic/gin"

func Success(c *gin.Context, code int, msg string, res any) {
	c.JSON(code, gin.H{
		"msg":      msg,
		"response": res,
	})
}

func Error(c *gin.Context, code int, msg string, err error) {
	c.JSON(code, gin.H{
		"msg":   msg,
		"error": err.Error(),
	})
}
