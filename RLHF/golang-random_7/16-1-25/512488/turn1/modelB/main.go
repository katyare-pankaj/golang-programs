package main

import (
	"fmt"
)

type Product struct {
	Name  string
	Price float64
}

func main() {
	product1 := Product{Name: "Laptop", Price: 1000.00}
	product2 := Product{Name: "Mouse", Price: 25.00}
	products := []Product{product1, product2}

	htmlPage := generateHTMLPage(products)
	fmt.Println(htmlPage)
}

func generateHTMLPage(products []Product) string {
	html := `
	<html>
	<head>
		<title>Products</title>
	</head>
	<body>
		<h1>Products</h1>
		<ul>
			%s
		</ul>
	</body>
	</html>
	`
	var productItems string
	for _, product := range products {
		productItem := fmt.Sprintf("<li>%s - $%.2f</li>", product.Name, product.Price)
		productItems += productItem
	}
	return fmt.Sprintf(html, productItems)
}
