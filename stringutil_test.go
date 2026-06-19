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

func TestTruncate(t *testing.T) {
	cases := []struct {
		name string
		s    string
		n    int
		want string
	}{
		{"shorter", "abc", 5, "abc"},
		{"equal", "abc", 3, "abc"},
		{"longer", "abcdef", 3, "abc"},
		{"negative n", "abc", -1, ""},
		{"empty n zero", "", 0, ""},
		{"empty n positive", "", 2, ""},
		{"n zero nonempty", "abc", 0, ""},
		{"multibyte truncate", "héllo", 3, "hél"},
		{"multibyte n ge len", "héllo", 10, "héllo"},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			if got := Truncate(tc.s, tc.n); got != tc.want {
				t.Fatalf("Truncate(%q,%d)=%q, want %q", tc.s, tc.n, got, tc.want)
			}
		})
	}
}

func TestPad(t *testing.T) {
	cases := []struct {
		name string
		s    string
		n    int
		want string
	}{
		{"shorter", "abc", 5, "abc  "},
		{"equal", "abc", 3, "abc"},
		{"longer unchanged", "abcdef", 3, "abcdef"},
		{"n zero", "abc", 0, "abc"},
		{"n negative", "abc", -1, "abc"},
		{"empty n positive", "", 2, "  "},
		{"empty n zero", "", 0, ""},
		{"multibyte shorter", "héllo", 7, "héllo  "},
		{"multibyte equal", "héllo", 5, "héllo"},
		{"multibyte longer", "héllo", 3, "héllo"},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			if got := Pad(tc.s, tc.n); got != tc.want {
				t.Fatalf("Pad(%q,%d)=%q, want %q", tc.s, tc.n, got, tc.want)
			}
		})
	}
}
