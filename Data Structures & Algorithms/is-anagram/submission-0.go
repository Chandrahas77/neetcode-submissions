import "slices"
func isAnagram(s string, t string) bool {
    if len(s) != len(t){
        return false
    }
    sBytes := []byte(s)
    tBytes := []byte(t)
    slices.Sort(sBytes)
    slices.Sort(tBytes)
    return string(sBytes) == string(tBytes)
}
