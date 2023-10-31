package response

import (
	"github.com/gin-gonic/gin"
)

type Json map[string]interface{}

type Response struct {
	code int
	message string
}


func Success(data map[string]interface{}) gin.H {
	var ret = gin.H{
		"code" : 200,
		"message" : "Success",
	}

	for k, v := range data {
		ret[k] = v
	}

	return ret
}


func InternalServerError(data map[string]interface{}) gin.H {
	var ret = gin.H{
		"code" : 500,
		"message" : "InternalServerError",
	}

	for k, v := range data {
		ret[k] = v
	}

	return ret
}


func Unauthorized(data map[string]interface{}) gin.H {
	var ret = gin.H{
		"code" : 401,
		"message" : "Unauthorized",
	}

	for k, v := range data {
		ret[k] = v
	}

	return ret
}

func BadRequest(data map[string]interface{}) gin.H {
	var ret = gin.H{
		"code" : 400,
		"message" : "BadRequest",
	}

	for k, v := range data {
		ret[k] = v
	}

	return ret
}