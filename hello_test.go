package hello_test

import (
	"fmt"
	"testing"

	"github.com/LoveCatdd/hello"
)

func TestHello(t *testing.T) {
	data := "jack"
	expect := fmt.Sprintf("hello, %s!", data)
	result := hello.Hello(data)
	if expect != result {
		t.Fatalf("expected result %s, but got %s", expect, result)
	}
}
