package util

import (
	"github.com/gin-gonic/gin"
	"io"
	"os"
)

func Logger() {
	gin.DisableConsoleColor()
	f, _ := os.Create("log/kbf.log")
	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)
}
