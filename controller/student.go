package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-student-admin/dao"
	"github.com/go-student-admin/util"
)

func GetStuInfo(c *gin.Context) {
	stuID := util.ToInt64(c.Query("id"))
	stu := dao.GetStuInfo(stuID)
	c.JSON(http.StatusOK, stu)
	return
}

func NewStudent(c *gin.Context) {
	name := c.PostForm("name")
	age := util.ToInt8(c.PostForm("age"))
	sex := util.ToInt8(c.PostForm("sex"))
	c.JSON(http.StatusOK, dao.NewStudent(name, age, sex))
}
