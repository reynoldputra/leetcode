func longestCommonPrefix(strs []string) string {
    i:= 0
    store := ""

    for true {
        stop := false
        if i >= len(strs[0])  {
            break
        }

        char := strs[0][i]

        for j:= 1; j < len(strs); j++ {
            if i >= len(strs[j])  {
                stop = true
                break
            }

            if char != strs[j][i] {
                stop = true
                break
            }
        }

        if stop {
            break
        }

        store = store + string(char)
        i++
    }


    return store
}
