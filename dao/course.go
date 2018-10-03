package dao

var (
	scoreGradeSql = `
		update
			schedule
		set
			grade = ?
		where
			student_id = ?
			and course_id = ?
	`
	getMaxGradeStuSql = `
		select
			t2.id as student_id,
			name,
			age,
			sex,
			course_id,
			grade
		from
			schedule t1 join student_info t2
			on t1.course_id = ?
			and t1.student_id = t2.id
		order by grade desc
		limit 1
	`
)

type Course struct {
	Id      int64  `gorm:"column:id" json:"id"`
	Name    string `gorm:"column:name" json:"name"`
	Credit  int8   `gorm:"column:credit" json:"credit"`
	Teacher string `gorm:"column:teacher" json:"teacher"`
}

func (course *Course) TableName() string {
	return "course"
}

type CseCourse struct {
	StudentId int64 `gorm:"column:student_id" json:"student_id"`
	CourseId  int64 `gorm:"column:course_id" json:"course_id"`
	Grade     int8  `gorm:"column:grade" json:"grade"`
}

func (cseCourse *CseCourse) TableName() string {
	return "schedule"
}

type MaxGradeStu struct {
	StudentId int64  `gorm:"column:student_id" json:"student_id"`
	Name      string `gorm:"column:name" json:"name"`
	Age       int8   `gorm:"column:age" json:"age"`
	Sex       int8   `gorm:"column:sex" json:"sex"`
	CourseId  int64  `gorm:"column:course_id" json:"course_id"`
	Grade     int8   `gorm:"column:grade" json:"grade"`
}

func NewCourse(name string, credit int8, teacher string) (course *Course) {
	conn, err := DBSession.GetConnection()
	if err != nil {
		return
	}
	course = &Course{
		Name:    name,
		Credit:  credit,
		Teacher: teacher,
	}
	conn.Create(course)
	return
}

func ChooseCourse(studentId int64, courseId int64) (cseCourse *CseCourse) {
	conn, err := DBSession.GetConnection()
	if err != nil {
		return
	}
	cseCourse = &CseCourse{
		StudentId: studentId,
		CourseId:  courseId,
	}
	conn.Create(cseCourse)
	return
}

func ScoreGrade(studentId int64, courseId int64, grade int8) (cseCourse *CseCourse) {
	conn, err := DBSession.GetConnection()
	if err != nil {
		return
	}
	var score CseCourse
	conn.Raw(scoreGradeSql, grade, studentId, courseId).Scan(score)
	cseCourse = &score
	cseCourse.Grade = grade
	return
}

func GetMaxGradeStu(courseId int64) (student *MaxGradeStu) {
	conn, err := DBSession.GetConnection()
	if err != nil {
		return
	}
	var stu MaxGradeStu
	conn.Raw(getMaxGradeStuSql, courseId).Scan(&stu)
	student = &stu
	return
}
