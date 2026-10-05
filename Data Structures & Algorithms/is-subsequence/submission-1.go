func isSubsequence(s string, t string) bool {
    if s == ""{
        return true
    }
    i := 0
    for j := range t{
        if s[i] == t[j]{
            i++
            if len(s) == i{
                return true
            }    
        }
    }
    return false
}
