package find_common_characters

func commonChars(words []string) []string {
	var res []string

	freq := make([][26]int, len(words))
	for i := range words {
		for j := 0; j < len(words[i]); j++ {
			freq[i][words[i][j]-'a'] += 1
		}
	}

	for char := 0; char < 26; char++ {
		minFreq := freq[0][char]

		for i := 1; i < len(words); i++ {
			if freq[i][char] == 0 {
				minFreq = 0
				break
			}

			if freq[i][char] < minFreq {
				minFreq = freq[i][char]
			}
		}

		for i := 0; i < minFreq; i++ {
			res = append(res, string('a'+byte(char)))
		}
	}

	return res
}
