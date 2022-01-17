package main

import (
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/gin-gonic/gin"
)

func main() {
	s := gin.Default()
	s.GET("/", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"msg": "Hello world~",
		})
	})

	go func() {
		err := s.Run(":8080")
    if err != nil {
			log.Fatalf("cannot start server: %v\n", err)
		}
	}()

	quitServer()
}

// 优雅退出服务
func quitServer() {
	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	// TODO: 服务退出时，处理各种操作，如释放资源
}
