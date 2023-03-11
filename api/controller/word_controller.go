package controller

import (
	"github.com/gin-gonic/gin"
)

// @Summary word search
// @Schemes
// @Description search word
// @Accept json
// @Produce json
// @Success 200 {object} model.WordNode
// @Router /api/word/search [get]
func WordSearch(g *gin.Context) {
	g.JSON(200, "aaaa")
}
