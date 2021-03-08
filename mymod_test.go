package mymod

import "testing"

func TestFoo(t *testing.T) {
    want := "foo"
    if got := Foo(); got != want {
        t.Errorf("Foo() = %q, want %q", got, want)
    }
}
