package main

import (
	"fmt"
	"testing"
)

////////////////////////////////////////////////////////////////////////////
// Lists

//==========================================================================
// list0

func TestList0(t *testing.T) {
	t.Log("First and plainest list test")
	const await = "The colors are: red, blue, white, .\n"
	if got := Generate(false, "list0"); got != await {
		t.Errorf("Mismatch:, got '%s' instead", got)
	}
}

//==========================================================================
// list1

func TestList1Text(t *testing.T) {
	t.Log("Second test, with text template")
	const await = "The quoted colors are: \"red\", \"blue\", \"white\", .\n"
	if got := Generate(false, "list1"); got != await {
		t.Errorf("Mismatch:, got '%s' instead", got)
	}
}

func TestList1HTML(t *testing.T) {
	t.Log("Second test, with html template")
	const await = "The quoted colors are: &#34;red&#34;, &#34;blue&#34;, &#34;white&#34;, .\n"
	if got := Generate(true, "list1"); got != await {
		t.Errorf("Mismatch:, got '%s' instead", got)
	}
}

//==========================================================================
// listfunc1

func TestListFunc1(t *testing.T) {
	t.Log("Test custom template function - minus1")
	const await = "red, blue, white"
	if got := Generate(false, "listfunc1"); got != await {
		t.Errorf("Mismatch:, got '%s' instead", got)
	}
}

func ExampleFunc1() {
	fmt.Print(Generate(false, "listfunc1"))
	// Output:
	// red, blue, white
}

////////////////////////////////////////////////////////////////////////////
// Commandline definitions

func ExampleTestExample() {
	fmt.Println(`  flags.Bool("debug", false, "Turn on debugging.")`)
	fmt.Println(`  viper.BindPFlag("debug", flags.Lookup("debug"))`)
	// Output:
	//   flags.Bool("debug", false, "Turn on debugging.")
	//   viper.BindPFlag("debug", flags.Lookup("debug"))
}

func ExampleCommandLineCobraViper() {
	fmt.Print(Generate(false, "commandlineCV"))
	// EasyGen commandlineCV | sed 's|^\t|&//&|; s|^$|\t//|'
	// Output:
	//
	//	flags.Bool("debug", false, "Turn on debugging.")
	//	viper.BindPFlag("debug", flags.Lookup("debug"))
	//
	//	flags.String("addr", "localhost:5002", "Address of the service.")
	//	viper.BindPFlag("addr", flags.Lookup("addr"))
	//
	//	flags.String("smtp-addr", "localhost:25", "Address of the SMTP server.")
	//	viper.BindPFlag("smtp-addr", flags.Lookup("smtp-addr"))
	//
	//	flags.String("smtp-user", "", "User for the SMTP server.")
	//	viper.BindPFlag("smtp-user", flags.Lookup("smtp-user"))
	//
	//	flags.String("smtp-password", "", "Password for the SMTP server.")
	//	viper.BindPFlag("smtp-password", flags.Lookup("smtp-password"))
	//
	//	flags.String("email-from", "noreply@abc.com", "The from email address.")
	//	viper.BindPFlag("email-from", flags.Lookup("email-from"))
	//
}
