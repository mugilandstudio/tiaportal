package server

import "github.com/gin-gonic/gin"

type FrontendServer struct{}

func NewFrontendServer() *FrontendServer {
	return &FrontendServer{}
}

func (s *FrontendServer) Start() error {
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
	})
	r.Run()
	return nil
}
