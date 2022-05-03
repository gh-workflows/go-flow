package wordgame


// IsPalindrome checks strings for matches
func IsPalindrome(word string) bool{

	for i := range word{
		l := len(word)
		if word[i] != word[l - i - 1]{
			return false
		}
	}
	return true
}