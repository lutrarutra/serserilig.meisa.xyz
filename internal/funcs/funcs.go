package funcs

import (
	"html/template"
)

var Functions = template.FuncMap{
	"add": add,
}

func add(num1, num2 int) int {
	return num1 + num2
}
