package wordgame


// IsPalindrome checks strings 
func IsPalindrome(word string) bool{

	for i := range word{
		l := len(word)
		if word[i] != word[l - i - 1]{
			return false
		}
	}
	return true
}