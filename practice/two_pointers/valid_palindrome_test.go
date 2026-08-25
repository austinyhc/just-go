package twopointers

import "testing"

func TestIsPalindrome(t *testing.T) {
	cases := []struct {
		name string
		s    string
		want bool
	}{
		{"leetcode example 1", "A man, a plan, a canal: Panama", true},
		{"leetcode example 2", "race a car", false},
		{"empty after stripping", " ", true},
		{"single char", "a", true},
		{"mixed case punctuation", "Was it a car or a cat I saw?", true},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			if got := IsPalindrome(c.s); got != c.want {
				t.Errorf("IsPalindrome(%q) = %v, want %v", c.s, got, c.want)
			}
		})
	}
}
