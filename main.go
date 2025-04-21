package main
import ("os")

func main() {
	lines, err := readLinesFromFile("main.java.txt")
	if err != nil {
		panic(err)
	}

	htmlContent := generateHTML(lines)
	os.WriteFile("output.html", []byte(htmlContent), 0644)
}
