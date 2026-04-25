package greatest_common_divisor_of_strings

import "strings"

func gcdOfStrings(str1 string, str2 string) string {
	if len(str2) > len(str1) {
		tmp := str1
		str1 = str2
		str2 = tmp
	}

	if !strings.HasPrefix(str1, str2) {
		return ""
	}

	prefix := str2

	for {
		if str1 == strings.Repeat(prefix, int(len(str1)/len(prefix))) && str2 == strings.Repeat(prefix, int(len(str2)/len(prefix))) {
			return prefix
		}

		if len(prefix) > 0 {
			prefix = prefix[:len(prefix)-1]
		}

		if len(prefix) == 0 {
			break
		}
	}

	return prefix
}
