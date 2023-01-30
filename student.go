package student

type Student struct {
	studentId   int
	studentName string
	gpa         float32
}

func NewStudent(studentId int, studentName string, gpa float32) *Student {
	return &Student{studentId: studentId, studentName: studentName, gpa: gpa}
}

func (s *Student) StudentId() int {
	return s.studentId
}

func (s *Student) SetStudentId(studentId int) {
	s.studentId = studentId
}

func (s *Student) StudentName() string {
	return s.studentName
}

func (s *Student) SetStudentName(studentName string) {
	s.studentName = studentName
}

func (s *Student) Gpa() float32 {
	return s.gpa
}

func (s *Student) SetGpa(gpa float32) {
	s.gpa = gpa
}
