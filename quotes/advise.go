package quotes

import (
	"fmt"
)

// Get_words_of_wisdom returns our recommended advice
func Get_words_of_wisdom() string {
	message := fmt.Sprintf("Quote: %s\nAuthor: %s",
		"Unless you try to do something beyond what you have already mastered, you will never grow.",
		"Ralph Waldo Emeron")
	return message
}
