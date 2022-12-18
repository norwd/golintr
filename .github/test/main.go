package main

import (
  "fmt"
  _ "pkg/src"
)

var ExportedVariableWithNoComment = "Hello World"

// this is not a vaild godoc comment
var ExportedVariableWithIncorrectComment = ExportedVariableWithNoComment

var _variableWithUnderscore = ExportedVariableWithNoComment

func main() {
  fmt.Println(ExportedVariableWithNoComment)
}
