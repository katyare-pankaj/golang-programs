package main

import (
	"fmt"
)

type Student struct {
	Name    string
	Grade   float64
	Comment string
}

func main() {
	students := []Student{
		{Name: "Alice", Grade: 95.5, Comment: "Excellent"},
		{Name: "Bob", Grade: 88.2, Comment: "Good"},
		{Name: "Charlie", Grade: 76.0, Comment: "Fair"},
	}

	htmlPage := generateStudentHTMLPage(students)
	fmt.Println(htmlPage)
}

func generateStudentHTMLPage(students []Student) string {
	htmlTemplate := `
	<html>
	<head>
		<title>Student Grades</title>
		<style>
			table {
				width: 100%;
				border-collapse: collapse;
			}
			th, td {
				border: 1px solid #ddd;
				padding: 8px;
				text-align: left;
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
				<th>Name</th>
				<th>Grade</th>
				<th>Comment</th>
			</tr>
			%s
		</table>
	</body>
	</html>
	`

	var studentRows string
	for _, student := range students {
		studentRow := fmt.Sprintf(
			`<tr>
				<td>%s</td>
				<td>%.2f</td>
				<td>%s</td>
			</tr>`,
			student.Name,
			student.Grade,
			student.Comment,
		)
		studentRows += studentRow
	}

	return fmt.Sprintf(htmlTemplate, studentRows)
}
