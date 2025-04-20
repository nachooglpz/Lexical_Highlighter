package main

import (
	"fmt"
	"html"
	"strings"
)

// Reference: Help from this stack overflow post:
// https://stackoverflow.com/questions/63124689/how-to-generate-html-in-go-with-data-from-struct
func tokenToHTML(token Token) string {
	escaped := html.EscapeString(token.Value)
	return fmt.Sprintf("<span class=\"%s\">%s</span>", token.Type, escaped)
}

// Reference: Help from this stack overflow post:
// https://stackoverflow.com/questions/1760757/how-to-efficiently-concatenate-strings-in-go
func generateHTML(lines []string) string {
	// Initialize the HTML content with a header and style
	var htmlContent strings.Builder

    htmlContent.WriteString("<html>\n<head>\n<style>\n")
    htmlContent.WriteString(`
        .keyword { color: blue; }
        .string { color: green; }
        .char { color: red; }
        .number { color: purple; }
        .operator { color: orange; }
        .punctuation { color: brown; }
        .whitespace { color: gray; }
        .identifier { color: black; }
        .comment { color: darkgreen; }
        .unknown { color: gray; } 
    `)
    htmlContent.WriteString("</style>\n</head>\n<body>\n")

    // Build body
    for _, line := range lines {
        tokens := getTokensFromLine(line)
        for _, token := range tokens {
            htmlContent.WriteString(tokenToHTML(token))
        }
        htmlContent.WriteString("<br>\n") // New line for each line of code
    }

    // Close the HTML tags
    htmlContent.WriteString("</body>\n</html>")
    
    return htmlContent.String()
}
