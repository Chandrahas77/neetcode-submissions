func isAnagram(s string, t string) bool {
    rs, rt := []rune(s), []rune(t)
    if len(rs) != len(rt){
        return false
    }
    counts := make(map[rune]int)
    for _,r := range rs{
        counts[r]++
    }
    for _,r := range rt{
        counts[r]--
        if counts[r] < 0{
            return false
        }
    }
    return true
}
