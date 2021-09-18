package lex

import (
	"bytes"
	"fmt"
	"strings"
)

type ModelP struct {
	bytes.Reader
	// result map[string]string
}

func (m ModelP) Read(s string) string {

	fmt.Println(strings.TrimRight(s, "EndModel"))
	return strings.TrimRight(s, "EndModel")
}

func (m ModelP) Result() {

}
