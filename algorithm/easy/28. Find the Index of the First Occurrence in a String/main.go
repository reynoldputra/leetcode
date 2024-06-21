func strstr(haystack string, needle string) int {
    for i,_ := range(haystack) {
        j := 0
        for j < len(needle) {
            if i + j >= len(haystack) {
                break
            }
            if needle[j] != haystack[i+j] {
                break
            }
            j++
        }

        if j >= len(needle) {
            return i
        }
    }
    
    return -1
}
