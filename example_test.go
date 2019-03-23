package easygen_test

import (
	"os"

	"github.com/go-easygen/easygen"
	"github.com/go-easygen/easygen/egVar"
)

// for standalone test, change package to `main` and the next func def to,
// func main() {
func ExampleProcess() {
	tmpl0 := easygen.NewTemplate().Customize()
	tmpl := tmpl0.Funcs(easygen.FuncDefs()).Funcs(egVar.FuncDefs())
	tmplFileName := "test/var0"
	easygen.Process(tmpl, os.Stdout, tmplFileName)
	easygen.Process2(tmpl, os.Stdout, tmplFileName, tmplFileName)

	// To use Execute(), TemplateFileName has to be exact
	m := easygen.ReadDataFile(tmplFileName + ".yaml")
	easygen.Execute(tmpl, os.Stdout, tmplFileName+".tmpl", m)

	// Output:
	// Input: "some-init-method"
	// Output 1: "SomeInitMethod"
	// Output 2: "SOME_INIT_METHOD"
	// Input: "some-init-method"
	// Output 1: "SomeInitMethod"
	// Output 2: "SOME_INIT_METHOD"
	// Input: "some-init-method"
	// Output 1: "SomeInitMethod"
	// Output 2: "SOME_INIT_METHOD"
}

// To show the full code in GoDoc
type dummy struct {
}
