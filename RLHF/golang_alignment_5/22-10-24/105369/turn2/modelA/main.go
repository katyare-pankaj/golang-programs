package main

import (
	"fmt"
	"strings"

	"github.com/alecthomas/chroma/formatters/html"
	"github.com/alecthomas/chroma/styles"
)

func measureIndentationAndVisualize(code string) string {
	indentationLevels := []int{}
	lines := strings.Split(code, "\n")

	for _, line := range lines {
		trimmedLine := strings.TrimLeft(line, " \t")
		indentationLevel := len(line) - len(trimmedLine)
		indentationLevels = append(indentationLevels, indentationLevel)
	}

	// Visualization: Generate HTML with syntax highlighting and indentation markers

	style := styles.Get("monokai")
	formatter := html.New(html.TabWidth(2))

	htmlCode := formatter.Format(nil, style, nil)

	return htmlCode.Error()
}

func main() {
	codeExample := `
		package main
		import (
			"fmt"
		)

		func main() {
			fmt.Println("Hello, World!")
			if true {
				fmt.Println("Indented block")
				if false {
					fmt.Println("Even more indented")
				}
			}
		}
	`

	visualizedCode := measureIndentationAndVisualize(codeExample)
	fmt.Println("Visualized Code:")
	fmt.Println(visualizedCode)
}
