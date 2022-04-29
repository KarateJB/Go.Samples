package greetings

import (
	"regexp"
	"testing"
)

func TestHelloName(t *testing.T) {
	name := "JB"
	want := regexp.MustCompile(`\b` + name + `\b`) // Compile the regular expression rule
	msg, err := Hello("JB")
	if !want.MatchString(msg) || err != nil { // If the result is not expected
		t.Fatalf(`Hello("JB") = %q, %v, want match for %#q, nil`, msg, err, want)
	}
}

func TestHelloEmpty(t *testing.T) {
	msg, err := Hello("")
	if msg != "" || err == nil { // If the result is not expected
		t.Fatalf(`Hello("") = %q, %v, want "", error`, msg, err)
	}
}
