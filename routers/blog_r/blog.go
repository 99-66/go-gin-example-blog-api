package blog_r

import (
	"github.com/99-66/go-gin-example-blog-api/controller/blog_c"
	"github.com/gin-gonic/gin"
)

func InitBlogRoutes(r *gin.Engine) {
	//r.POST("/login")
	//r.GET("/logout")

	api := r.Group("/api")
	{
		v1 := api.Group("/1")
		{
			user := v1.Group("/user")
			{
				user.POST("", blog_c.CreateUser)
				user.GET("/:id", blog_c.FindUserById)
				user.PUT("/:id", blog_c.UpdateUser)
				user.DELETE("/:id", blog_c.DeleteUser)
			}
		}
	}
}
