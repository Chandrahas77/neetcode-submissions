func groupAnagrams(strs []string) [][]string {
	res := make(map[string][]string)

	for _,s := range strs{
		sortedStr := sortString(s)
		res[sortedStr] = append(res[sortedStr],s)
	}

	var result [][]string

	for _,group := range res{
		result = append(result,group)
	}
	return result
}

func sortString(s string) string{
	characters := []rune(s)
	sort.Slice(characters,func(i,j int) bool{
		return characters[i] < characters[j]
	})
	return string(characters)
}
