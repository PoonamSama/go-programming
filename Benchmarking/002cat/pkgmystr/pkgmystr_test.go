package mystr

import (
	"fmt"
	"strings"
	"testing"
)

func TestCatenate(t *testing.T) {
	s := "Shaken not stirred"
	xs := strings.Split(s, " ")
	s = Catenate(xs)
	if s != "Shaken not stirred" {
		t.Error("got:", s, "want:", "Shaken not stirred")
	}
}
func TestJoin(t *testing.T) {
	s := "Shaken not stirred"
	xs := strings.Split(s, " ")
	s = Join(xs)
	if s != "Shaken not stirred" {
		t.Error("got:", s, "want:", "Shaken not stirred")
	}
}
func ExampleCatenate() {
	s := "Shaken not stirred"
	xs := strings.Split(s, " ")
	fmt.Println(Catenate(xs))
	// Output:
	// Shaken not stirred
}
func ExampleJoin() {
	s := "Shaken not stirred"
	xs := strings.Split(s, " ")
	fmt.Println(Catenate(xs))
	// Output:
	// Shaken not stirred
}

const s = "A constant string that has many words,therefore creating a sentence. The sentence in the present tense,may not make any sense but don't be so tense"

var xs []string

func BenchmarkCatenate(b *testing.B) {
	xs = strings.Split(s, " ")
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Catenate(xs)
	}
}

func BenchmarkJoin(b *testing.B) {
	xs = strings.Split(s, " ")
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Join(xs)
	}
}
