package server

import (
	"fmt"
	"html/template"
	"net/http"
	"path/filepath"

	"github.com/AYehia0/lezer/utils"
	"github.com/gin-gonic/gin"
)

var tmpPath string

// share/serve the files in a specific directory
func (server *Server) serveFiles(ctx *gin.Context) {

	// get the path from the url
	currentDir, _ := ctx.Params.Get("dir")
	var currentAbs string

	if currentDir == "/" {
		server.CurrentDir, _ = filepath.Abs(server.ServerDir)
		server.ServerDir = server.CurrentDir
		currentAbs = server.CurrentDir
	} else {
		currentAbs, _ = filepath.Abs(filepath.Join(server.ServerDir, currentDir))
	}

	// 1. get all the files and dirs in a specific dir
	items, err := utils.GetAllInDir(server.ServerDir, currentAbs, server.IgnoreHidden)
	if err != nil {
		ctx.JSON(http.StatusNoContent, fmt.Sprintf("%s", err))
		return
	}
	// 2. return these files to the client
	ctx.HTML(http.StatusOK, "list.html", gin.H{
		"files":        items,
		"currentDir":   template.URL(server.ServerDir),
		"relativePath": template.URL(filepath.Clean(currentDir)),
	})
}
