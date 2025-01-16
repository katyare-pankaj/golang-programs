package main

import (
	"fmt"
)

type Student struct {
	Name  string
	Grade int
}

func generateStudentGradeHTML(students []Student) string {
	htmlTemplate := `
    <html>
    <head>
        <title>Student Grades</title>
        <style>
            table {
                width: 80%;
                margin: 20px auto;
                border-collapse: collapse;
            }
            th, td {
                padding: 10px;
                border: 1px solid #ddd;
            }
            th {
                background-color: #f2f2f2;
            }
        </style>
    </head>
    <body>
        <h1>Student Grades</h1>
        <table>
            <tr>
                <th>Student</th>
                <th>Grade</th>
            </tr>
            %s
        </table>
    </body>
    </html>    
    `
	var studentRows string
	for _, student := range students {
		studentRow := fmt.Sprintf(`
        <tr>
            <td>%s</td>
            <td>%d</td>
        </tr>
        `, student.Name, student.Grade)
		studentRows += studentRow
	}
	return fmt.Sprintf(htmlTemplate, studentRows)
}

func main() {
	students := []Student{
		{Name: "Alice", Grade: 88},
		{Name: "Bob", Grade: 79},
		{Name: "Charlie", Grade: 92},
	}
	htmlPage := generateStudentGradeHTML(students)
	fmt.Println(htmlPage)
}
