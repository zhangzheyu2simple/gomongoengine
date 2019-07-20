package main

import "testing"

func TestPrint(t *testing.T) {
	r := print()
	if r != 1 {
		t.Errorf("err")
	}
}
