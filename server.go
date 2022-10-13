package main

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Run() {
	r := gin.Default()
	r.Static("/download", *dir)
	err := r.Run(fmt.Sprintf(":%d", *port))
	if err != nil && !errors.Is(err, http.ErrServerClosed) {
		logger.Errorf("run err: %s", err)
	}
}
