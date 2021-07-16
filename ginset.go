package ginset

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func New(paths []string, assetFunc func(string) ([]byte, error), system http.FileSystem) gin.HandlerFunc {
	return func(c *gin.Context) {
		if c.Request.Method != "" && c.Request.Method != "GET" {
			c.Next()
			return
		}
		for _, v := range paths {
			if c.Request.URL.Path != v {
				continue
			}
			url := strings.TrimLeft(strings.TrimRight(c.Request.URL.Path, "/"), "/") + "/"
			if url == "/" {
				url = ""
			}
			b, err := assetFunc(url + "index.html")
			if err != nil {
				c.AbortWithError(http.StatusInternalServerError, err)
				return
			}
			c.Data(http.StatusOK, "text/html", b)
			c.Next()
			return
		}
		c.FileFromFS(c.Request.URL.Path, system)
		c.Next()
	}
}
