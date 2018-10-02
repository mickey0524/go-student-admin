package main

import (
	"github.com/go-student-admin/dao"
	"github.com/go-student-admin/route"
)

func main() {
	dao.Init()
	route.Init()
}
