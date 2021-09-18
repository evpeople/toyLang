package lex

import (
	"fmt"
	"testing"
)

func TestRead(t *testing.T) {
	code := "Model dsadasd aEndModel"
	var a ModelP
	b := a.Read(code)
	fmt.Println(b)
}
