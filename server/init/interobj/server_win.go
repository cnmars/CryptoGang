//go:build windows
// +build windows

package interobj

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func InitServer(address string, router *gin.Engine) server {
	return &http.Server{
		Addr:           address,
		Handler:        router,
		ReadTimeout:    20 * time.Second,
		WriteTimeout:   20 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
}
