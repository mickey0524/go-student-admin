package dao

var (
	stuInfoSql = `
		select
			id,
			name,
			age,
			sex
		from
			student_info
		where
			id = ?
	`
)

type Student struct {
	Id   int64  `gorm:"column:id" json:"id"`
	Name string `gorm:"column:name" json:"name"`
	Age  int8   `gorm:"column:age" json:"age"`
	Sex  int8   `gorm:"column:sex" json:"sex"`
}

func GetStuInfo(id int64) (stuInfo *Student) {
	conn, err := DBSession.GetConnection()
	if err != nil {
		return
	}
	var stu Student
	conn.Raw(stuInfoSql, id).Scan(&stu)
	stuInfo = &stu
	return
}

func NewStudent(name string, age int8, sex int8) (stuInfo *Student) {
	conn, err := DBSession.GetConnection()
	if err != nil {
		return
	}
	stuInfo = &Student{
		Name: name,
		Age:  age,
		Sex:  sex,
	}
	conn.Create(stuInfo)
	return
}

func (student *Student) TableName() string {
	return "student_info"
}
