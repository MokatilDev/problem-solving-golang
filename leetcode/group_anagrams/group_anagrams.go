package group_anagrams

import (
	"slices"
)

func groupAnagrams(strs []string) [][]string {
	var res [][]string
	groupes := make(map[string][]string)

	for _, s := range strs {
		runnes := []rune(s)
		slices.Sort(runnes)

		key := string(runnes)

		groupes[key] = append(groupes[key], s)
	}

	for _, groupe := range groupes {
		res = append(res, groupe)
	}

	return res
}
