package egCal_test

import (
	"os"

	"github.com/go-easygen/easygen"
	"github.com/go-easygen/easygen/egCal"
	"github.com/go-easygen/easygen/egVar"
)

// for standalone test, change package to `main` and the next func def to,
// func main() {
func Example() {
	tmpl0 := easygen.NewTemplate().Customize()
	tmpl := tmpl0.Funcs(easygen.FuncDefs()).Funcs(egVar.FuncDefs()).Funcs(egCal.FuncDefs())
	err := easygen.Process0(tmpl, os.Stdout,
		"{{.Name}}: {{clk2uc .Name}} {{clk2ss .Name}}\n"+
			"Cal: {{add 2 3}}, {{multiply 2 3}}, {{subtract 9 2}}, {{divide 24 3}}\n",
		"../test/var0")

	// Output:
	// some-init-method: SomeInitMethod SOME_INIT_METHOD
	// Cal: 5, 6, 7, 8

	if err != nil {
		panic(err)
	}
}

// To show the full code in GoDoc
type dummy struct {
}
