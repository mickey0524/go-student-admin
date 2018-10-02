package controller

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-student-admin/dao"
)

func GetStuInfo(c *gin.Context) {
	stuID, _ := strconv.ParseInt(c.Query("id"), 10, 64)
	stu := dao.GetStuInfo(stuID)
	c.JSON(http.StatusOK, stu)
	return
}
