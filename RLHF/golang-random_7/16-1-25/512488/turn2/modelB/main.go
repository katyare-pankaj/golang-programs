func generateHTMLPage(products []Product) string {
	html := `<html>
            <head>
                <title>Products</title>
            </head>
            <body>
                <h1>Products</h1>
                <ul>
                    %s
                </ul>
            </body>
            </html>`
	var productList string
	for _, product := range products {
		productList += fmt.Sprintf(`<li><h3>%s</h3><p>Price: $%.2f</p></li>`, product.Name, product.Price)
	}
	return fmt.Sprintf(html, productList)
}