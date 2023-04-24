package main

import (
	"os"
	"text/template"
)

// test.tpl = "base"
// a/test.tpl = "a"
// a_test.tpl = "a"
// b/test.tpl = "b"
// b_test.tpl = "b"

func main() {
	tpl, _ := template.ParseFiles("test.tpl", "a/test.tpl", "b/test.tpl")
	tpl.Execute(os.Stdout, nil)
	tpl, _ = template.ParseFiles("test.tpl", "a_test.tpl", "b_test.tpl")
	tpl.Execute(os.Stdout, nil)
}
