package handlers

import (
	"dfunani/mydecoder/models"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

const pagesBasePath = "/pages"

func HandleHomePage(c *gin.Context) {
	tab := "encode"
	if strings.HasSuffix(strings.TrimSuffix(c.Request.URL.Path, "/"), "/decode") {
		tab = "decode"
	}
	c.HTML(http.StatusOK, "home.html", models.HomePage{
		Title:     "MyDecoder",
		Tagline:   "Turn a JSON object into a compact token — and back again.",
		PagesBase: pagesBasePath,
		ActiveTab: tab,
	})
}
