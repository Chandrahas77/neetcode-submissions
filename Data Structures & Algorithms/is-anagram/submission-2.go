func isAnagram(s string, t string) bool {
    if len(s) != len(t){
        return false
    }
    var counts [26]int
    for i := 0; i < len(s); i++{
        sIdx := s[i] - 'a'
        tIdx := t[i] - 'a'
        counts[sIdx]++
        counts[tIdx]--
    }
    for _,val := range counts{
        if val!=0 {
            return false
        }
    }
    return true
}
