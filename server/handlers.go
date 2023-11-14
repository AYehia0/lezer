package server

import (
	"fmt"
	"html/template"
	"net/http"
	"path/filepath"

	"github.com/AYehia0/lezer/utils"
	"github.com/gin-gonic/gin"
)

// handle the main route : the index page
func (server *Server) handleMainIndex(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "main.html", gin.H{
		"BrowseRoute": template.URL(server.Config.BrowseRoute),
	})
}

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

	// check if the path is a file
	if utils.IsDownloadable(currentAbs) {
		// download the file
		ctx.File(currentAbs)
		return
	}

	items, err := utils.GetAllInDir(server.ServerDir, currentAbs, server.IgnoreHidden)
	if err != nil {
		ctx.JSON(http.StatusNoContent, fmt.Sprintf("Path not found: %s  \nError: %s", currentAbs, err))
		return
	}
	ctx.HTML(http.StatusOK, "list.html", gin.H{
		"files":        items,
		"currentDir":   template.URL(server.ServerDir),
		"relativePath": template.URL(filepath.Clean(currentDir)),
		"BrowseRoute":  template.URL(server.Config.BrowseRoute),
	})
}

// upload files from android to linux
func (server *Server) uploadFiles(ctx *gin.Context) {
	file, err := ctx.FormFile("file")
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}
	err = ctx.SaveUploadedFile(file, filepath.Join(server.DownloadLocation, file.Filename))

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	ctx.HTML(http.StatusOK, "upload_success.html", gin.H{"url": template.URL(server.DownloadLocation)})
}
