package main

import (
	"fmt"

	"github.com/proogramadorJ/monkey-interpreter/lexer"
	"github.com/proogramadorJ/monkey-interpreter/token"
)

func main() {
	var tk token.Token
	var l lexer.Lexer

	tk.Type = "STRING"
	fmt.Println(tk.Type)
	//fmt.Println(l.position)

}
