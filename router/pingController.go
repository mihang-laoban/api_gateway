package router

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

func Ping(c *gin.Context) {
	c.JSON(200, gin.H{
		"msg": "pong",
	})
	return
}

func PingThis(c *gin.Context) {
	//json.Marshal()
	//http.Post("http://localhost/this", "application/json; charset=utf-8", )

	str := `{
		"data": "this"
	}`
	data := strings.NewReader(str)
	res, err := http.Post("http://localhost:5000/this", "application/json; charset=utf-8", data)
	if err != nil || res.StatusCode != http.StatusOK {
		c.Status(http.StatusServiceUnavailable)
		return
	}

	reader := res.Body
	contentLength := res.ContentLength
	contentType := res.Header.Get("Content-Type")

	extraHeaders := map[string]string{
		//"Content-Disposition": `attachment; filename="gopher.png"`,
	}

	c.DataFromReader(http.StatusOK, contentLength, contentType, reader, extraHeaders)
}
