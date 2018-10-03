package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-student-admin/dao"
	"github.com/go-student-admin/util"
)

func NewCourse(c *gin.Context) {
	name := c.PostForm("name")
	credit := util.ToInt8(c.PostForm("credit"))
	teacher := c.PostForm("teacher")
	c.JSON(http.StatusOK, dao.NewCourse(name, credit, teacher))
}

func ChooseCourse(c *gin.Context) {
	studentId := util.ToInt64(c.PostForm("student_id"))
	courseId := util.ToInt64(c.PostForm("course_id"))
	c.JSON(http.StatusOK, dao.ChooseCourse(studentId, courseId))
}

func ScoreGrade(c *gin.Context) {
	studentId := util.ToInt64(c.PostForm("student_id"))
	courseId := util.ToInt64(c.PostForm("course_id"))
	grade := util.ToInt8(c.PostForm("grade"))
	c.JSON(http.StatusOK, dao.ScoreGrade(studentId, courseId, grade))
}

func GetMaxGradeStu(c *gin.Context) {
	course_id := util.ToInt64(c.Query("course_id"))
	c.JSON(http.StatusOK, dao.GetMaxGradeStu(course_id))
}
