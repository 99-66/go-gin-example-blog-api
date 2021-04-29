package blog_c

import (
	"errors"
	"github.com/99-66/go-gin-example-blog-api/libs"
	"github.com/99-66/go-gin-example-blog-api/models/blog_m"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
	"strconv"
)

func CreateUser(c *gin.Context) {
	var user blog_m.User
	err := c.ShouldBindJSON(&user)
	if err != nil {
		libs.ErrResponse(c, http.StatusBadRequest, "Bad Request")
		return
	}

	user, err = user.Create()
	if err != nil {
		libs.ErrResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code": http.StatusOK,
		"message": user,
	})
}

func FindUserById(c *gin.Context) {
	p := c.Param("id")
	id, err := strconv.Atoi(p)
	if err != nil {
		libs.ErrResponse(c, http.StatusBadRequest, "Bad Request")
		return
	}

	var user blog_m.User
	user, err = user.FindUserById(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			libs.ErrResponse(c, http.StatusNotFound, "Not Found")
			return
		}

		libs.ErrResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code": http.StatusOK,
		"message": user,
	})
}

func UpdateUser(c *gin.Context) {
	p := c.Param("id")
	id, err := strconv.Atoi(p)
	if err != nil {
		libs.ErrResponse(c, http.StatusBadRequest, "Bad Request")
		return
	}

	var user blog_m.User
	user, err = user.FindUserById(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			libs.ErrResponse(c, http.StatusNotFound, "Not Found")
			return
		}

		libs.ErrResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	var updateUser blog_m.User
	err = c.ShouldBindJSON(&updateUser)
	if err != nil {
		libs.ErrResponse(c, http.StatusBadRequest, "Bad Request")
		return
	}

	err = user.UpdateUser(&updateUser)
	if err != nil {
		libs.ErrResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code": http.StatusOK,
		"message": updateUser,
	})
}

func DeleteUser(c *gin.Context) {
	p := c.Param("id")
	id, err := strconv.Atoi(p)
	if err != nil {
		libs.ErrResponse(c, http.StatusBadRequest, "Bad Request")
		return
	}

	var user blog_m.User
	user, err = user.FindUserById(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			libs.ErrResponse(c, http.StatusNotFound, "Not Found")
			return
		}

		libs.ErrResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	err = user.DeleteUserById(id)
	if err != nil {
		libs.ErrResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusNoContent, nil)
}