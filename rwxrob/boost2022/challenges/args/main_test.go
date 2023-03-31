package main

import "testing"

func TestOut(t *testing.T) {
	out := output()
	if out != "this" {
		t.Errorf("\nwant: %q\ngot:  %q\n", "this", out)

	}
}
