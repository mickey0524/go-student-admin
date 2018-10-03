package route

import (
	"github.com/gin-gonic/gin"
	"github.com/go-student-admin/controller"
)

func Init() {
	router := gin.Default()

	router.GET("/stu_info", controller.GetStuInfo)
	router.GET("/max_grade_stu", controller.GetMaxGradeStu)

	router.POST("/new_stu", controller.NewStudent)
	router.POST("/new_course", controller.NewCourse)
	router.POST("/choose_course", controller.ChooseCourse)
	router.POST("/score_grade", controller.ScoreGrade)

	router.Run()
}
