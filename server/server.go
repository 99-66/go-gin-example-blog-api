package server

import (
	"github.com/99-66/go-gin-example-blog-api/routers/blog_r"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func RunAPI(address string) error {
	return RunAPIWithMiddleware(address)
}

func RunAPIWithMiddleware(address string) error {
	r := gin.Default()

	// Set CORS Middlewares
	corsConf := cors.DefaultConfig()
	corsConf.AllowAllOrigins = true
	r.Use(cors.New(corsConf))

	// Set Routes
	blog_r.InitBlogRoutes(r)

	return r.Run(address)
}