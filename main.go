package student

import (
	"fmt"
)

func main() {
	fmt.Println("Hello World!")
	student := NewStudent(1, "Nuradil", 3.92)

	fmt.Printf("Name of the student %v\n", student.StudentName())
	fmt.Printf("Id of the student %v\n", student.StudentId())
}
