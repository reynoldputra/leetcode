func lengthOfLastWord(s string) int {
    for i := len(s) - 1 ; i >= 0 ; i-- {
        if s[i] != 32 {
            for j := i ; j >= 0 ; j-- {
                if s[j] == 32 {
                    fmt.Println(i,j)
                    return i - j
                }
            }
            return i + 1
        }
    }

    return 0
}
