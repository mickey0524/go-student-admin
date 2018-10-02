package route

import (
	"github.com/gin-gonic/gin"
	"github.com/go-student-admin/controller"
)

func Init() {
	router := gin.Default()

	router.GET('/stu_info', controller.GetStuInfo)

	router.Run()
}
