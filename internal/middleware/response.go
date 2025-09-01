package middleware

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"

	"github.com/gin-gonic/gin"
)

type writer struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

func (w *writer) Write(data []byte) (int, error) {
	w.body.Write(data)
	return w.ResponseWriter.Write(data)
}

type ResponseMiddleware struct{}

func NewResponseMiddleware() *ResponseMiddleware {
	return &ResponseMiddleware{}
}

func (rm *ResponseMiddleware) Handler() gin.HandlerFunc {
	return func(c *gin.Context) {
		w := &writer{
			ResponseWriter: c.Writer,
			body:           &bytes.Buffer{},
		}
		c.Writer = w
		c.Next()

		status := c.Writer.Status()

		if len(c.Errors) > 0 {
			err := c.Errors[0].Error()
			fmt.Println(err)

			if err == "" {
				err = "Internal Error"
			}

			fmt.Println(err)
			c.JSON(status, gin.H{
				"error": err,
			})
		} else {
			if json.Valid(w.body.Bytes()) {
				var originalData any
				json.Unmarshal(w.body.Bytes(), &originalData)
				c.JSON(status, gin.H{"data": originalData})
			} else {
				io.Copy(c.Writer, w.body)
			}

		}
	}
}
