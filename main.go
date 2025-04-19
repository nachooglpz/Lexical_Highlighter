package main

import "fmt"

func main() {
	fmt.Println(getTokensFromLine("String hello = \"Hello, World!\"; // This is a comment"));
	fmt.Println()
	fmt.Println(getTokensFromLine("int x = 42;"));
	fmt.Println()
	fmt.Println(getTokensFromLine("if (x > 0) { return x; }"));
	fmt.Println()
	fmt.Println(getTokensFromLine("char c = 'a';"));
	fmt.Println()
	fmt.Println(getTokensFromLine("boolean flag = true;"));
	fmt.Println()
	fmt.Println(getTokensFromLine("for (int i = 0; i < 10; i++) { System.out.println(i); }"));
}
