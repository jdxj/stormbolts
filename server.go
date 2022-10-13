package main

import (
	"errors"
	"fmt"
	"net/http"
	"path/filepath"

	"github.com/gin-gonic/gin"
)

func Run() {
	r := gin.Default()
	register(r)
	err := r.Run(fmt.Sprintf(":%d", *port))
	if err != nil && !errors.Is(err, http.ErrServerClosed) {
		logger.Errorf("run err: %s", err)
	}
}

func register(r gin.IRouter) {
	r.POST("/upload", upload)
}

func upload(ctx *gin.Context) {
	fh, err := ctx.FormFile("my-file")
	if err != nil {
		logger.Errorf("FormFile err: %s\n", err)
		ctx.JSON(http.StatusBadRequest, gin.H{
			"msg": "get file err",
		})
		return
	}

	path := filepath.Join(*dir, fh.Filename)
	err = ctx.SaveUploadedFile(fh, path)
	if err != nil {
		logger.Errorf("SaveUploadedFile err: %s", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"msg": "save file err",
		})
		return
	}

	ctx.JSON(http.StatusOK, nil)
}

func download(ctx *gin.Context) {}
