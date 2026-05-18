import "slices"

func groupAnagrams(strs []string) [][]string {
	groups := make(map[string][]string)
	for _, str := range strs {
		b := []byte(str)
		slices.Sort(b)
		key := string(b)
		groups[key] = append(groups[key], str)
	}

	var target [][]string
	for _, value := range groups {
		target = append(target, value)
	}

	return target
}
