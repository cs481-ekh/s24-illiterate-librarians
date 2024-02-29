// serve the static content of the site
package static

import (
	"github.com/gin-gonic/gin"
)

// StaticHandler handles the static content of the site
func StaticHandler(c *gin.Context) {
	c.File("/client/index.html")
}
