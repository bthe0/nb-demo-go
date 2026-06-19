package stringutil

import "testing"

func TestReverse(t *testing.T) {
	if got := Reverse("abc"); got != "cba" {
		t.Fatalf("Reverse(abc) = %q, want cba", got)
	}
}

func TestWordCount(t *testing.T) {
	if got := WordCount("a b  c"); got != 3 {
		t.Fatalf("WordCount = %d, want 3", got)
	}
}
