package saying

import (
	"fmt"
	"testing"
)

func TestWelcome(t *testing.T) {
	s := Welcome("James")
	if s != "hello dearJames" {
		t.Error("expected hello dearJames but got", s)
	}
}
func ExampleWelcome() {
	fmt.Println(Welcome("James"))
	//Output:
	//hello dearJames
}
func BenchmarkWelcome(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Welcome("James")
	}
}
