package greetings

import (
	"regexp"
	"testing"
)

func TestHelloName(t *testing.T) {
	//arrange
	name := "Test"
	want := regexp.MustCompile(`\b` + name + `\b`)

	//ack
	msg, err := Hello(name)

	//assert
	if !want.MatchString(msg) || err != nil {
		t.Fatalf(`Hello("Gladys") = %q, %v, want match for %#q, nil`, msg, err, want)
	}
}

func TestHelloEmpty(t *testing.T) {
	//arrange
	name := ""

	//ack
	msg, err := Hello(name)

	//assert
	if msg != "" || err == nil {
		t.Fatalf(`Hello("") = %q, %v, want "", error`, msg, err)
	}
}

func TestHellos(t *testing.T) {
	//arrange
	names := []string{"A", "B", "C"}

	//ack
	msgs, err := Hellos(names)

	//assert
	if len(msgs) == 0 || err != nil {
		t.Fatalf(`Hellos({"A", "B", "C"}) %v, want %q, error`, msgs, err)
	}
}

func TestHellosEmpty(t *testing.T) {
	//arrange
	names := []string{"A", "B", ""}

	//ack
	msgs, err := Hellos(names)

	//assert
	if len(msgs) > 0 || err == nil {
		t.Fatalf(`Hellos({"A", "B", ""}) %v, want %q, error`, msgs, err)
	}
}
