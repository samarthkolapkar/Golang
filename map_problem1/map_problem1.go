/*Write a Go program to store the grades of students in a map. The program should allow the user to:

Add a new student along with their grade.
Update the grade of an existing student.
Delete a student from the record.
Print the list of all students along with their grades.
Your program should use a map with student names as keys and grades as values. Grades can be represented as integers.
*/
package main

import "fmt"

func main() {
	student := make(map[string]int)
	var number_of_students int
	fmt.Print("Enter the number of students:")
	fmt.Scan(&number_of_students)
	for i := 0; i < number_of_students; i++ {
		var name string
		var grades int
		fmt.Println("Enter the name of student:")
		fmt.Scan(&name)
		fmt.Println("Enter the garde of student:")
		fmt.Scan(&grades)

		student[name] = grades
	}
	fmt.Println(student)

}
