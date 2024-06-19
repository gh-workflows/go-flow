package wordgame

import (
	"testing"
)

func TestPalindome(t *testing.T) {
	var words = []string{"kayak", "abba", "121", "bob"}
	for _, word := range words {
		if !IsPalindrome(word) {
			t.Errorf("'%s' was incorrectly not identifed as palindrome", word)
		}
	}

}

func TestNonPalindome(t *testing.T) {
	var words = []string{"apple", "orange", "pear"}
	for _, word := range words {
		if IsPalindrome(word) {
			t.Errorf("'%s' was incorrectly identifed as palindrome", word)
		}
	}
}
