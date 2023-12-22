func maximumWealth(accounts [][]int) int {
    max := 0
    for _, arr := range accounts {
        w := 0
        for _, n := range arr {
            w += n
        }
        if w > max {
            max = w
        }
    }

    return max
}
