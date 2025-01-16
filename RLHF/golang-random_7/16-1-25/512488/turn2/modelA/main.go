package main

import (
	"fmt"
)

type Product struct {
	Name        string
	Price       float64
	Description string
}

func main() {
	products := []Product{
		{Name: "Laptop", Price: 999.99, Description: "A high-performance laptop with a 15-inch screen."},
		{Name: "Smartphone", Price: 699.99, Description: "A flagship smartphone with a 6.5-inch display."},
		{Name: "Headphones", Price: 199.99, Description: "Noise-canceling headphones with premium sound."},
	}

	htmlPage := generateHTMLPage(products)
	fmt.Println(htmlPage)
}

func generateHTMLPage(products []Product) string {
	htmlTemplate := `
	<html>
	<head>
		<title>Product List</title>
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
		<h1>Product List</h1>
		<table>
			<tr>
				<th>Name</th>
				<th>Price</th>
				<th>Description</th>
			</tr>
			%s
		</table>
	</body>
	</html>
	`

	var productRows string
	for _, product := range products {
		productRow := fmt.Sprintf(
			`<tr>
				<td>%s</td>
				<td>$%.2f</td>
				<td>%s</td>
			</tr>`,
			product.Name,
			product.Price,
			product.Description,
		)
		productRows += productRow
	}

	return fmt.Sprintf(htmlTemplate, productRows)
}
