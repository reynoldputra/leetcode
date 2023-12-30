func maxProfit(prices []int) int {
    r := 10007
    p := 0

    for _, v := range prices {
        if v < r {
            r = v
        } else if v - r > p {
            p = v - r
        }
    }

    return p
}
