package main

import "fmt"

var ExportedVariableWithNoComment = "Hello World"

// this is not a vaild godoc comment
var ExportedVariableWithIncorrectComment = ExportedVariableWithNoComment

func main() {
  fmt.Println(ExportedVariableWithNoComment)
}
